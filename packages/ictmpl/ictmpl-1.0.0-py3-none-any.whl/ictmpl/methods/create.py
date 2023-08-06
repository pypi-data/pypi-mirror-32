import os
from re import compile
from os.path import abspath, isdir, join
from json import loads
from subprocess import Popen, PIPE, check_call
from uuid import uuid4
from textwrap import dedent

from ictmpl.helpers.git import Git
from ictmpl.helpers.fsutil import (
    copytree_with_ignore, rmtree_without_ignore, replace_template_file)
        

__all__ = ('create_project',)


RE_LOCALDIR = compile(r'^(?:\.{1,2}/|/)')
RE_GITLINK = compile(r'^(?:git@|https?://)')

RE_COMMAND_ENV = compile(r'#!(.+)')


def create_project(args, app):
    project_path = abspath(args.path)
    
    tmpl = args.template
    if RE_LOCALDIR.match(tmpl):
        template_path = abspath(tmpl)
        if not isdir(template_path):
            raise Exception("Template wasn't found")
        
        copytree_with_ignore(template_path, project_path)
    elif RE_GITLINK.match(tmpl):
        Git(tmpl).clone_to(project_path)
    else:
        root_repository = app.getconf('REPOSITORY', None)
        if not root_repository:
            raise Exception("Template wasn't found")
        
        Git(root_repository + tmpl).clone_to(project_path)
    
    rmtree_without_ignore(project_path)
    os.chdir(project_path)
    
    rc = {}
    try:
        with open('.ictmplrc', 'r') as file:
            exec(file.read(), rc)
    except:
        return
    
    dependencies = rc.get('sysDependencies', None)
    if dependencies and isinstance(dependencies, dict):
        commands = []
        for key, packages in dependencies.items():
            pipes = Popen('command -v %s' % key, shell=True, stdout=PIPE)
            stdout, _ = pipes.communicate()
            if not stdout:
                commands.append(packages)
    
        if commands:
            # TODO: Need check platform
            command = 'sudo apt-get install %s' % ' '.join(commands)
            Popen(command, shell=True)
    
    for key in ('name', 'version', 'description', 'author', 'author_email',
                'author_website', 'licence'):
        app.params['__{}__'.format(key.upper())] = rc.get(key, '')
    
    params = rc.get('params', {})
    if params:
        print('Configure template:')
        for param, default in params.items():
            question = 'SET {}: ({}) '.format(param, default)
            app.params[param] = input(question) or default
        print('------------------\n')
    
        print('Walking files and replace params')
        templates = rc.get('templates', [])
        if templates:
            templates = map(lambda p: join(project_path, p), templates)
            for filepath in templates:
                replace_template_file(filepath, app)
        else:
            for root, dirs, files in os.walk(project_path):
                for filename in files:
                    replace_template_file(join(root, filename), app)
    
    setup_command = rc.get('commands', {}).get('setup', None)
    if setup_command:
        if isinstance(setup_command, (list, tuple)):
            setup_command = '\n'.join(setup_command)
        
        print('Run setup command')
        setup_command = dedent(setup_command).strip()
        env = RE_COMMAND_ENV.search(setup_command)
        if not env:
            env = '/bin/bash'
        
        env = env.group(1).strip()
        tmpfilename = '/tmp/ictmpl_%s' % uuid4().hex
        try:
            check_call('echo "{}" > {}'.format(setup_command, tmpfilename),
                       shell=True, stdout=None)
            check_call('{} {}'.format(env, tmpfilename), shell=True)
            check_call('rm -rf {}'.format(tmpfilename), shell=True)
        except KeyboardInterrupt:
            print('KeyboardInterrupt')
