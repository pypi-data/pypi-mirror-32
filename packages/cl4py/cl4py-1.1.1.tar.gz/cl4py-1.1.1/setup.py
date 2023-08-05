"""The cl4py setup.

See:
https://github.com/marcoheisig/cl4py
"""

from setuptools import setup, find_packages
from codecs import open
from os import path

def readme():
    with open('README.rst') as f:
        return f.read()

setup(
    name='cl4py',
    version='1.1.1',
    description='Common Lisp for Python',
    long_description=readme(),
    license='MIT',
    url='https://github.com/marcoheisig/cl4py',
    author='Marco Heisig',
    author_email='marco.heisig@fau.de',
    packages=find_packages(exclude=['contrib', 'docs', 'test']),
    install_requires=[],
    extras_require={},
    keywords='foreign functions FFI',
    classifiers=[
        'Development Status :: 3 - Alpha',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
        'Programming Language :: Lisp' ,
    ],
    package_data={'cl4py': ['cl4py/py.lisp']},
)
