#!/usr/bin/env python
# -*- coding: utf-8 -*-
from setuptools import setup, find_packages

with open("README.md", "r") as fh:
    long_description = fh.read()

setup(
    name='mqtt_client',
    version='1.0.1',
    description='Simple MQTT Client.',
    long_description=long_description,
    long_description_content_type="text/markdown",
    keywords='mqtt',
    author='samuel.deancos',
    author_email='sdeancos@gmail.com',
    packages=find_packages(),
    include_package_data=False,
    scripts=['bin/mqtt_client'],
    python_requires='>=3.6',
    install_requires=[
        'docopt',
        'paho-mqtt',
        'terminaltables'
    ]
)
