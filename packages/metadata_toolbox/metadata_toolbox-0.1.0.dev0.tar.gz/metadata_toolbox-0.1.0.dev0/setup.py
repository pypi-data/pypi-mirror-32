#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from setuptools import setup, find_packages

PROJECT = 'Toolbox to Manage Corpora & Their Metadata'
VERSION = '0.1'
REVISION = '0.1.0.dev0'
AUTHOR = 'Students of SE Software-Projekte, WS 2017/2018'

setup(
    name='metadata_toolbox',
    version=REVISION,
    description=PROJECT,
    # url
    author=AUTHOR,
    # license
    classifiers=[
        'Development Status :: 3 - Alpha',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6'
    ],
    # keywords
    packages=find_packages(exclude=['docs', 'test']),
    install_requires=[
        'parse>=1.8.2',
        'pandas>=0.19.2',
        'lxml>=3.6.4'
    ],
    command_options={
        'build_sphinx': {
            'project': ('setup.py', PROJECT),
            'version': ('setup.py', VERSION),
            'release': ('setup.py', REVISION),
        }
    }
)
