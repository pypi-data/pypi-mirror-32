#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import sys

from setuptools import setup, find_packages

if sys.argv[-1] == 'publish':
    try:
        import wheel

        print("Wheel version: ", wheel.__version__)
    except ImportError:
        print('Wheel library missing. Please run "pip install wheel"')
        sys.exit()
    os.system('python setup.py sdist upload')
    os.system('python setup.py bdist_wheel upload')
    sys.exit()

readme = open('README.rst').read()
history = open('HISTORY.rst').read().replace('.. :changelog:', '')

setup(
    name='django-surveys',
    version='1.0.0',
    description="""A reusable Django app that lets users write feedback for any model""",
    long_description=readme + '\n\n' + history,
    author='Founders4Schools',
    author_email='dev@founders4schools.org.uk',
    url='https://github.com/founders4schools/django-surveys',
    packages=find_packages('src'),
    package_dir={'': 'src'},
    include_package_data=True,
    install_requires=[
        "django",
        "django-compat>=1.0.8",
        "django-model-utils",
        "djangorestframework>=3.2",
        "django-braces>=1.10.0",
    ],
    license="MIT",
    zip_safe=False,
    keywords='django-surveys',
    classifiers=[
        'Development Status :: 3 - Alpha',
        'Framework :: Django',
        'Framework :: Django :: 1.11',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: BSD License',
        'Natural Language :: English',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
    ],
)
