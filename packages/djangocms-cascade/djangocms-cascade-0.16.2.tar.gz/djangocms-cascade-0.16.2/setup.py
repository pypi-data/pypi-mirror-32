#!/usr/bin/env python
# -*- coding: utf-8 -*-
from __future__ import unicode_literals

import io
from setuptools import setup, find_packages
from cmsplugin_cascade import __version__

def readfile(filename):
    with io.open(filename, encoding='utf-8') as fd:
        return fd.read()

CLASSIFIERS = [
    'Development Status :: 5 - Production/Stable',
    'Environment :: Web Environment',
    'Framework :: Django :: 1.9',
    'Framework :: Django :: 1.10',
    'Intended Audience :: Developers',
    'License :: OSI Approved :: MIT License',
    'Operating System :: OS Independent',
    'Programming Language :: Python',
    'Topic :: Internet :: WWW/HTTP :: Dynamic Content',
    'Programming Language :: Python :: 2.7',
    'Programming Language :: Python :: 3.4',
    'Programming Language :: Python :: 3.5',
    'Programming Language :: Python :: 3.6',
]

setup(
    name='djangocms-cascade',
    version=__version__,
    description='Collection of extendible plugins for django-CMS to create and edit widgets in a simple manner',
    author='Jacob Rief',
    author_email='jacob.rief@gmail.com',
    url='https://github.com/jrief/djangocms-cascade',
    packages=find_packages(exclude=['examples', 'docs', 'tests']),
    install_requires=[
        'jsonfield',
        'django-classy-tags>=0.8.0',
        'django-cms>=3.4.5',
        'djangocms-text-ckeditor>=3.4.0',
    ],
    license='LICENSE-MIT',
    platforms=['OS Independent'],
    classifiers=CLASSIFIERS,
    long_description=readfile('README.md'),
    long_description_content_type='text/markdown',
    include_package_data=True,
    zip_safe=False,
)
