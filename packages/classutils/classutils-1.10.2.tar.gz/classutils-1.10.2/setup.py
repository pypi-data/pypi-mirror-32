# encoding: utf-8

import os
import re
from setuptools import setup, find_packages


MODULE_NAME = u'classutils'  # The module name must match this!

metadata_file = open(os.path.join(MODULE_NAME, u'_metadata.py')).read()
metadata = dict(re.findall("__([a-z]+)__\s*=\s*u'([^']+)'", metadata_file))

test_dependencies = []


def read(filename):

    """
    Utility function used to read the README file into the long_description.

    :param filename: Filename to read

    :return: file pointer
    """

    return open(os.path.join(os.path.dirname(__file__), filename)).read()


setup(
    name=MODULE_NAME,

    # Versions should comply with PEP440.  For a discussion on single-sourcing
    # the version across setup.py and the project code, see
    # https://packaging.python.org/en/latest/single_source_version.html
    version=metadata.get(u'version'),

    author=metadata.get(u'author'),
    author_email=metadata.get(u'authoremail'),

    license=metadata.get(u'license'),

    url=u'https://{host}/{user}/{repo}'.format(host=metadata.get(u'githost'),
                                               user=metadata.get(u'gituser'),
                                               repo=metadata.get(u'gitrepo')),
    download_url=u'https://{host}/{user}/{repo}/get/{version}.tar'.format(host=metadata.get(u'githost'),
                                                                          user=metadata.get(u'gituser'),
                                                                          repo=metadata.get(u'gitrepo'),
                                                                          version=metadata.get(u'version')),

    packages=find_packages(),

    # If you want to distribute just a my_module.py, uncomment
    # this and comment out packages:
    #   py_modules=["my_module"],

    description=metadata.get(u'description'),
    long_description=read(u'README.rst'),

    keywords=[
        u'Class',
        u'Observer'],

    # See https://pypi.python.org/pypi?%3Aaction=list_classifiers
    classifiers=[
        # How mature is this project?
        #   Development Status :: 1 - Planning
        #   Development Status :: 2 - Pre-Alpha
        #   Development Status :: 3 - Alpha
        #   Development Status :: 4 - Beta
        #   Development Status :: 5 - Production/Stable
        #   Development Status :: 6 - Mature
        #   Development Status :: 7 - Inactive
        u'Development Status :: 3 - Alpha',

        # Indicate who your project is intended for
        u'Intended Audience :: Developers',

        # Pick your license as you wish (should match "license" above)
        u'License :: OSI Approved :: MIT License',

        # Specify the Python versions you support here. In particular, ensure
        # that you indicate whether you support Python 2, Python 3 or both.
        u'Programming Language :: Python :: 2',
        u'Programming Language :: Python :: 2.7',
        u'Programming Language :: Python :: 3',
        u'Programming Language :: Python :: 3.6',

        u'Topic :: Utilities',
    ],

    # Dependencies
    install_requires=[
        u'pip>=9.0.1',
        u'future>=0.16.0',
        u'fdutil>=1.6.2',
        u'logging_helper>=1.5.0',
        u'timingsutil>=1.0.9',
        u'requests>=2.7.0',
        u'tableutil>=2.2.4',
        u'cachingutil>=1.0.18',
    ],

    tests_require=test_dependencies,

    extras_require={
        u'test': test_dependencies
    },

    # Reference any non-python files to be included here
    package_data={
        '': ['*.md', '*.rst', '*.db', '*.txt'],  # Include all files from any package that contains *.db/*.md/*.txt
        },

    # To provide executable scripts, use entry points in preference to the
    # "scripts" keyword. Entry points provide cross-platform support and allow
    # pip to create the appropriate form of executable for the target platform.
    entry_points={
        'console_scripts': [],
        },

    # List any scripts that are to be deployed to the python scripts folder
    scripts=[]

)
