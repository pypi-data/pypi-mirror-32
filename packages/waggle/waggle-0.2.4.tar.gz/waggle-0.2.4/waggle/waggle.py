import json
import os
import re
import subprocess
import sys


from botocore.exceptions import ClientError
import boto3

from argparse import ArgumentParser


def find_tasks(dirnames):
    "Find all the Docker directories that are subdirectories of the given list"
    tasks = []
    for dirname in dirnames:
        for path, subdirs, files in os.walk(dirname):
            if 'Dockerfile' in files:
                tasks.append(path)
    return tasks


def dependency(namespace, dirname):
    "Extract the dependency of the task described in a given directory"
    with open(os.path.abspath(os.path.join(dirname, 'Dockerfile'))) as dfile:
        from_lines = [
            line[5:].strip()
            for line in dfile
            if line.startswith('FROM ')
        ]
        if len(from_lines) == 1:
            if from_lines[0].startswith(namespace + '/'):
                return from_lines[0]

        return None


def order_by_dependency(namespace, dirnames):
    "Examine all the directories in `dirnames`, and order by dependency"

    dependencies = [
        (dirname, dependency(namespace, dirname))
        for dirname in dirnames
    ]
    tasks = set('%s/%s' % (namespace, os.path.basename(d)) for d in dirnames)

    # Now sort the tasks to ensure that dependencies are met. This
    # is done by repeatedly iterating over the input list of tasks.
    # If the dependency of a given task is on the final list,
    # that task is promoted to the end of the final list. This process
    # continues until the input list is empty, or we do a full iteration
    # over the input models without promoting a model to the final list.
    # If we do a full iteration without a promotion, that means there are
    # circular dependencies in the list.
    ordered = []
    resolved = set()
    while dependencies:
        skipped = []
        changed = False
        while dependencies:
            dirname, dep = dependencies.pop()
            task = '%s/%s' % (namespace, os.path.basename(dirname))

            # If the dependency is already on the final model list, or
            # not on the original serialization list, then we've found
            # another model with a satisfied dependency.
            if dep not in tasks or dep in resolved:
                resolved.add(task)
                ordered.append(dirname)
                changed = True
            else:
                skipped.append((dirname, dep))
        if not changed:
            print(
                '\n'
                "Can't resolve dependencies for tasks in %s." %
                ', '.join(dirname for dirname, dep in sorted(skipped))
            )
            sys.exit(13)
        dependencies = skipped

    return ordered


def register(namespace, tag, region, aws_access_key_id, aws_secret_access_key, *dirnames):
    bad_names = []
    not_a_dir = []
    tasks = []
    for dirname in dirnames:
        full_path = os.path.abspath(dirname)
        task_name = os.path.basename(full_path)
        if re.match('[^-_A-Za-z0-9]', task_name):
            bad_names.append(task_name)
        elif not os.path.isdir(full_path):
            not_a_dir.append(task_name)
        else:
            tasks.append((task_name, full_path))

    if bad_names:
        print()
        print("The following tasks will not be valid ECS tasks:", file=sys.stderr)
        for name in bad_names:
            print("    *", name, file=sys.stderr)

    if not_a_dir:
        print()
        print("The following paths don't appear to be Docker configurations:", file=sys.stderr)
        for name in not_a_dir:
            print("    *", name, file=sys.stderr)

    if bad_names or not_a_dir:
        return bad_names + not_a_dir

    print("Waggling %s..." % ', '.join(dirnames))

    print("Logging into ECR...")
    proc = subprocess.Popen([
            'aws', 'ecr', 'get-login',
            '--no-include-email',
            '--region', region
        ],
        stdout=subprocess.PIPE
    )
    result, errors = proc.communicate()

    proc = subprocess.Popen(
        result.decode('utf-8').strip().split(' '),
        cwd=full_path,
    )
    proc.wait()

    registered = []
    for task_name, full_path in tasks:
        _register(full_path, namespace, task_name, tag, region, aws_access_key_id, aws_secret_access_key)
        registered.append(task_name)

    print()
    print("The following tasks have been registered with AWS ECS:")
    for task in registered:
        print('    * %s' % task)


def _register(full_path, namespace, task_name, tag, region, aws_access_key_id, aws_secret_access_key):
    try:
        repository_name = "%s/%s" % (
            namespace,
            task_name
        )
        print("Registering %s as an AWS ECS task..." % repository_name)
        print("Building local Docker image for %s..." % repository_name)
        proc = subprocess.Popen([
                'docker', 'build',
                "-t", repository_name,
                '.'
            ],
            cwd=full_path,
        )
        proc.wait()

        print("Looking up AWS ECR repository URI...")
        aws_session = boto3.session.Session(
            region_name=region,
            aws_access_key_id=aws_access_key_id,
            aws_secret_access_key=aws_secret_access_key,
        )
        ecr = aws_session.client('ecr')
        try:
            response = ecr.describe_repositories(repositoryNames=[repository_name])
            uri = response['repositories'][0]['repositoryUri']
        except ClientError as e:
            response = ecr.create_repository(repositoryName=repository_name)
            uri = response['repository']['repositoryUri']
        print("   ECR repository URI is", uri)

        print("Tagging Docker image for publication...")
        proc = subprocess.Popen([
                'docker', 'tag',
                "%s:%s" % (repository_name, tag),
                "%s:%s" % (uri, tag),
            ],
            cwd=full_path,
        )
        proc.wait()

        print("Pushing Docker image to AWS...")
        proc = subprocess.Popen([
                'docker', 'push',
                "%s:%s" % (uri, tag),
            ],
            cwd=full_path,
        )
        proc.wait()

        print("Registering ECS task...")
        definition = {
            'name': task_name,
            'image': uri,
            'memory': 128,
            'cpu': 0,
            'logConfiguration': {
                'logDriver': 'awslogs',
                'options': {
                    'awslogs-group': namespace,
                    'awslogs-region': region,
                    'awslogs-stream-prefix': task_name
                }
            }
        }

        try:
            with open(os.path.join(full_path, 'ecs.json')) as data:
                config = json.load(data)
                print("ECS configuration overrides:")
                for key, value in sorted(config.items()):
                    print("    %s: %s" % (key, value))
                definition.update(config)
        except IOError:
            print("Couldn't load ecs.json configuration; using defaults")

        ecs = aws_session.client('ecs')
        response = ecs.register_task_definition(
            family=task_name,
            networkMode='bridge',
            containerDefinitions=[definition]
        )

        revision = response['taskDefinition']['revision']

        print("AWS ECS task registration complete.")

        return "%s:%s" % (task_name, revision)

    except Exception as e:
        print("Unable to prepare AWS task %s" % full_path)
        print(e)


def main():
    parser = ArgumentParser()
    parser.add_argument(
        '--tag', '-t', default='latest', dest='tag',
        help='Specify the version tag to use.',
    )
    parser.add_argument(
        '--namespace', default='beekeeper', dest='namespace',
        help='Specify the namespace to use for AWS services.',
    )
    parser.add_argument(
        'dirnames', metavar='dirname', nargs='+',
        help='Directories containing Docker image to register.'
    )
    options = parser.parse_args()

    # Load sensitive environment variables from a .env file
    try:
        with open('.env') as envfile:
            for line in envfile:
                if line.strip() and not line.startswith('#'):
                    key, value = line.strip().split('=', 1)
                    os.environ.setdefault(key.strip(), value.strip())
    except FileNotFoundError:
        pass

    tasks = find_tasks(options.dirnames)
    ordered = order_by_dependency(options.namespace, tasks)

    if not tasks:
        print()
        print("Couldn't find any tasks in %s" % ', '.join(options.dirnames), file=sys.stderr)
        sys.exit(2)

    try:
        failed = register(
            options.namespace,
            options.tag,
            os.environ['AWS_REGION'],
            os.environ['AWS_ACCESS_KEY_ID'],
            os.environ['AWS_SECRET_ACCESS_KEY'],
            *ordered
        )

        if failed:
            sys.exit(3)
    except KeyError as e:
        print("AWS environment variable %s not found" % e)
        sys.exit(1)


if __name__ == '__main__':
    main()
