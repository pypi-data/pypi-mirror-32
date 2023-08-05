# Copyright (c) 2015 - 2016 EMC Corporation.
# All Rights Reserved.
#
#    Licensed under the Apache License, Version 2.0 (the "License"); you may
#    not use this file except in compliance with the License. You may obtain
#    a copy of the License at
#
#         http://www.apache.org/licenses/LICENSE-2.0
#
#    Unless required by applicable law or agreed to in writing, software
#    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
#    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
#    License for the specific language governing permissions and limitations
#    under the License.

""" Python distutils setup for siolib distribution """

import codecs
import os
import setuptools


def read(fname):
    here = os.path.dirname(__file__)
    with codecs.open(os.path.join(here, fname), encoding='utf-8') as f:
        return f.read()


setuptools.setup(
    name='siolib',
    version='2.0.3',
    description='ScaleIO API base library',
    url='https://github.com/codedellemc/python-scaleioclient',
    license='ASL 2.0',
    author='Cloudscaling (DellEMC)',
    long_description=read('README.rst'),
    packages=['siolib'],
    classifiers=[
        # Reference: http://pypi.python.org/pypi?%3Aaction=list_classifiers
        'Development Status :: 4 - Beta',
        'Intended Audience :: Developers',
        'Operating System :: OS Independent',
        'Programming Language :: Python',
        'Programming Language :: Python :: 2',
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3',
        'Topic :: Software Development :: Libraries :: Python Modules',
        'License :: OSI Approved :: Apache Software License'],
    install_requires=[
        'enum34',
        'requests',
        'urllib3',
        'six',
    ],
    keywords=[
        'emc',
        'scaleio'
    ],
)
