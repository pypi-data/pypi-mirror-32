#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""The setup script."""

try:
    from setuptools import setup, find_packages
except ImportError:
    from distutils.core import setup, find_packages

from mooquant_mootdx import __version__

with open('README.md') as readme_file:
    readme = readme_file.read()

requirements = ['mooquant', 'mootdx']
test_requirements = ['mooquant', 'mootdx', 'pytest', ]
setup_requirements = ['mooquant', 'mootdx']

setup(
    name='mooquant_mootdx',
    version=__version__,
    description="mooquant mootdx module",
    long_description=readme + '\n\n',
    author="bopo.wang",
    author_email='ibopo@126.com',
    url='https://github.com/bopo/mooquant_mootdx',
    packages=find_packages(include=['mooquant_mootdx', 'mooquant_mootdx.*']),
    include_package_data=True,
    install_requires=requirements,
    license="MIT license",
    zip_safe=False,
    keywords='mooquant_mootdx',
    entry_points={
        'console_scripts': [
            'mooquant_mootdx = mooquant_mootdx.cli:main',
        ]
    },
    classifiers=[
        'Development Status :: 2 - Pre-Alpha',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Natural Language :: English',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.3',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
    ],
    test_suite='tests',
    tests_require=test_requirements,
    setup_requires=setup_requirements,
)
