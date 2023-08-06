from __future__ import print_function

import os
import sys
import subprocess
import traceback

from setuptools import find_packages
from distutils.core import setup
from setuptools.command.install import install
import wheel  # pip install wheel -U
from wheel.bdist_wheel import bdist_wheel as bdist_wheel_
from multiprocessing import Process

PROJ_PATH = os.path.abspath(os.path.dirname(os.path.realpath(__file__)))
THIRD_PARTY = PROJ_PATH + '/pymagnitude/third_party'
PYSQLITE = THIRD_PARTY + '/_pysqlite'


def install_custom_sqlite3():
    # Begin install custom SQLite
    # Can be safely ignored even if it fails, however, system SQLite
    # limitations may prevent large .magnitude files with many columns
    # from working.
    print("Installing custom SQLite 3....")
    rc = subprocess.Popen([
        sys.executable,
        PYSQLITE + '/setup.py',
        'install',
        '--install-lib=' + THIRD_PARTY + '/internal',
    ], cwd=PYSQLITE).wait()
    if rc:
        print("")
        print("============================================================")
        print("=========================WARNING============================")
        print("============================================================")
        print("It seems like building a custom version of SQLite on your")
        print("machine has failed. This is fine, Magnitude will likely work")
        print("just fine with the sytem version of SQLite for most use cases.")
        print("However, if you are trying to load extremely high dimensional")
        print("models > 999 dimensions, you may run in to SQLite limitations")
        print("that can only be resolved by using the custom version of SQLite.")
        print("To troubleshoot make sure you have appropriate build tools on")
        print("your machine for building C programs like GCC and the standard")
        print("library. Also make sure you have the python-dev development")
        print("libraries and headers for building Python C extensions.")
        print("If you need more help with this, please reach out to ")
        print("opensource@plasticity.ai.")
        print("============================================================")
        print("============================================================")
        print("")
    else:
        print("")
        print("============================================================")
        print("=========================SUCCESS============================")
        print("============================================================")
        print("Building a custom version of SQLite on your machine has")
        print("succeeded.")
        print("============================================================")
        print("============================================================")
        print("")

    # End install custom SQLite


def copy_custom_sqlite3():
    # Copy the pysqlite2 folder into site-packages under
    # pymagnitude/third_party/internal/ for good measure
    try:
        import site
        from glob import glob
        from distutils.dir_util import copy_tree
        cp_from = THIRD_PARTY + '/internal/'
        for sitepack in site.getsitepackages():
            globbed = glob(sitepack + '/pymagnitude*/')
            try:
                cp_to = globbed[0] + '/pymagnitude/third_party/internal/'
            except IndexError as e:
                print(
                    "Site Package: '" +
                    sitepack +
                    "' did not have pymagnitude")
                continue
            print("Copying from: ", cp_from, " --> to: ", cp_to)
            copy_tree(cp_from, cp_to)
    except Exception as e:
        print("Error copying internal pysqlite folder:")
        traceback.print_exc(e)


class CustomBdistWheelCommand(bdist_wheel_):
    def run(self):
        install_custom_sqlite3()
        print("Running wheel...")
        p = Process(target=bdist_wheel_.run, args=(self,))
        p.start()
        p.join()
        print("Done running wheel...")
        copy_custom_sqlite3()


class CustomInstallCommand(install):
    def run(self):
        install_custom_sqlite3()
        print("Running install...")
        p = Process(target=install.do_egg_install, args=(self,))
        p.start()
        p.join()
        print("Done running install...")
        copy_custom_sqlite3()


if __name__ == '__main__':
    setup(
        name='pymagnitude',
        packages=find_packages(
            exclude=[
                'tests',
                'tests.*']),
        version='0.1.32',
        description='A fast, efficient universal vector embedding utility package.',
        long_description="""
    About
    -----
    A feature-packed Python package and vector storage file format for utilizing vector embeddings in machine learning models in a fast, efficient, and simple manner developed by `Plasticity <https://www.plasticity.ai/>`_. It is primarily intended to be a faster alternative to `Gensim <https://radimrehurek.com/gensim/>`_, but can be used as a generic key-vector store for domains outside NLP.

    Documentation
    -------------
    You can see the full documentation and README at the `GitLab repository <https://gitlab.com/Plasticity/magnitude>`_ or the `GitHub repository <https://github.com/plasticityai/magnitude>`_.
        """,
        author='Plasticity',
        author_email='opensource@plasticity.ai',
        url='https://gitlab.com/Plasticity/magnitude',
        keywords=[
                    'pymagnitude',
                    'magnitude',
                    'plasticity',
                    'nlp',
                    'natural',
                    'language',
                    'processing',
                    'word',
                    'vector',
                    'embeddings',
                    'embedding',
                    'word2vec',
                    'gensim',
                    'alternative',
                    'machine',
                    'learning',
                    'annoy',
                    'index',
                    'approximate',
                    'nearest',
                    'neighbors'],
        license='MIT',
        setup_requires=['numpy >= 1.14.0'],
        install_requires=[
            'pip >= 9.0.1',
            'numpy >= 1.14.0',
            'xxhash >= 1.0.1',
            'fasteners >= 0.14.1',
            'annoy >= 1.11.4',
            'lz4 >= 1.0.0'],
        classifiers=[
            "Development Status :: 5 - Production/Stable",
            'Intended Audience :: Developers',
            "Topic :: Software Development :: Libraries :: Python Modules",
            'Topic :: Scientific/Engineering :: Artificial Intelligence',
            'Topic :: Text Processing :: Linguistic',
            "Operating System :: OS Independent",
            'License :: OSI Approved :: MIT License',
            'Programming Language :: Python :: 2.7',
            'Programming Language :: Python :: 3',
            'Programming Language :: Python :: 3.0',
            'Programming Language :: Python :: 3.1',
            'Programming Language :: Python :: 3.2',
            'Programming Language :: Python :: 3.3',
            'Programming Language :: Python :: 3.4',
            'Programming Language :: Python :: 3.5',
            'Programming Language :: Python :: 3.6',
            'Programming Language :: Python :: 3.7'],
        cmdclass={
            'install': CustomInstallCommand,
            'bdist_wheel': CustomBdistWheelCommand
        },
    )
