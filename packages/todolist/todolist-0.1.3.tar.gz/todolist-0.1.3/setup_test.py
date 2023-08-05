#!/usr/bin/env python3'
from setuptools import setup, find_packages

setup(
    name             = 'todolist',
    version          = '0.1.3',
    description      = 'The Open Source CLI To Do List App',
    author           = 'Kang MinSeok',
    author_email     = 'stillnessfate@gmail.com',
    url              = 'https://github.com/StillnessFate/todo-list-cli',
    download_url     = 'https://github.com/StillnessFate/todo-list-cli/archive/master.tar.gz',
    install_requires = [ ],
    packages         = find_packages(exclude = ["test.py", "test_code.py"]),
    keywords         = ['todo list', 'task manager'],
    python_requires  = '>=3',
    package_data     =  {'todolist' : [ "LICENSE" ]},
    zip_safe=False,
    classifiers      = [
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.2',
        'Programming Language :: Python :: 3.3',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6'
    ],
    entry_points={
        'console_scripts': [
            'todolist = todolist.__main__:main'
        ]
    }
)