#!/usr/bin/env python

from setuptools import setup, find_packages


setup(name='atc-beta-helper',
      version='0.0.6',
      license='MIT',
      platforms='any',
      packages=find_packages(),
      install_requires=[
          'requests==2.18.4',
          'PyYAML==3.12'
      ],
      entry_points={
          "console_scripts": [
              "helper = autocnn_helper.config:write_parameter",
          ],
      },
      classifiers=[
          'Programming Language :: Python',
          'Operating System :: OS Independent',
      ])
