#!/usr/bin/env python
# -*- coding: utf-8 -*-
try:
    from setuptools import setup
except ImportError:
    from distutils.core import setup


requires = [
    'lxml',
    'requests',
    'Pillow'
]

setup(
    name='adbui',
    version='0.3.6',
    description='Python Wrapper for Android UiAutomator test tool',
    long_description='Python wrapper for Android uiautomator tool.',
    author='Tango Nian',
    author_email='hao1032@gmail.com',
    url='https://github.com/hao1032/adbui',
    keywords=[
        'testing', 'android', 'uiautomator', 'ocr'
    ],
    install_requires=requires,
    packages=['adbui'],
    license='MIT',
    platforms='any',
    classifiers=(
        'Development Status :: 4 - Beta',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python',
        'Programming Language :: Python :: 2',
        'Programming Language :: Python :: 3',
        'Topic :: Software Development :: Testing'
    )
)