# -*- coding: utf-8 -*-
# -----------------------------------------------------------------------------
# Copyright (c) Spyder Project Contributors
#
# Licensed under the terms of the MIT License
# (see LICENSE.txt for details)
# -----------------------------------------------------------------------------
"""Setup script for spyder_terminal."""

# Standard library imports
import ast
import os
import os.path as osp
import sys

# Third party imports
from setuptools import find_packages, setup

from setupbase import (BuildStatic, CleanComponents, COMPONENTS, HERE,
                       SdistWithBuildStatic)


def get_version(module='spyder_terminal'):
    """Get version."""
    with open(os.path.join(HERE, module, '__init__.py'), 'r') as f:
        data = f.read()
    lines = data.split('\n')
    for line in lines:
        if line.startswith('VERSION_INFO'):
            version_tuple = ast.literal_eval(line.split('=')[-1].strip())
            version = '.'.join(map(str, version_tuple))
            break
    return version


def get_description():
    """Get long description."""
    with open(os.path.join(HERE, 'README.rst'), 'r') as f:
        data = f.read()
    return data


REQUIREMENTS = ['spyder>=3.2.0', 'tornado',
                'coloredlogs', 'requests']


# Verify that COMPONENTS exist before trying to build wheels
if any([arg == 'bdist_wheel' for arg in sys.argv]):
    if not osp.isdir(COMPONENTS):
        print("\nWARNING: Server components are missing!! Please run "
              "'python setup.py sdist' first.\n")
        sys.exit(1)


# Add pywinpty to our Windows dependencies when building wheels
if os.name == 'nt' or any([arg.startswith('win') for arg in sys.argv]):
    REQUIREMENTS.append('pywinpty>=0.2.1')
else:
    REQUIREMENTS.append('pexpect')

cmdclass = {
    'build_static': BuildStatic,
    'sdist': SdistWithBuildStatic,
    'clean_components': CleanComponents
}


setup(
    name='spyder-terminal',
    version=get_version(),
    cmdclass=cmdclass,
    keywords=['Spyder', 'Plugin'],
    url='https://github.com/spyder-ide/spyder-terminal',
    license='MIT',
    author='Spyder Project Contributors',
    author_email='admin@spyder-ide.org',
    description='Spyder Plugin for displaying a virtual terminal '
                '(OS independent) inside the main Spyder window',
    long_description=get_description(),
    packages=find_packages(exclude=['contrib', 'docs', 'tests*']),
    install_requires=REQUIREMENTS,
    include_package_data=True,
    classifiers=[
        'Development Status :: 4 - Beta',
        'Intended Audience :: Developers',
        'Intended Audience :: Science/Research',
        'License :: OSI Approved :: MIT License',
        'Operating System :: MacOS',
        'Operating System :: Microsoft :: Windows',
        'Operating System :: POSIX :: Linux',
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6'
    ])
