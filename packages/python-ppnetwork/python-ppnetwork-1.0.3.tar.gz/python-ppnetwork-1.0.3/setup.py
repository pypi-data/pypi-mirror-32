#!/usr/bin/env python3
# coding=utf-8

from setuptools import setup, find_packages

setup(name='python-ppnetwork',
    py_modules = ["ppnetwork","pp_link","pp_vpn","pp_flow","pp_control"],
#       packages = find_packages(),
#       package_data = {"":["*.yaml'"]},
#       data_files=[('config', ['*.yaml']),],
#       exclude_package_data = { 'test': ['test*.py'] },
      entry_points = {
        'console_scripts': [
            'ppnetwork = ppnetwork:domain',
        ],},
      author='gonewind.he',
      author_email='gonewind.he@gmail.com',
      maintainer='gonewind',
      maintainer_email='gonewind.he@gmail.com',
      url='https://github.com/gonewind73/ppnetwork',
      description='Linux/Windows  Virtual LAN application written in  python',
      long_description=open('README.rst',encoding='utf-8').read(),
      version='1.0.3',
      install_requires=['python-pytuntap>=1.0.5',        ],
      python_requires='>=3',
      platforms=["Linux","Windows"],
      classifiers=[
          'Development Status :: 5 - Production/Stable',
          'Intended Audience :: Developers',
          'License :: OSI Approved :: MIT License',
          'Programming Language :: Python :: 3',
          'Topic :: Software Development :: Libraries :: Python Modules',
          'Topic :: System :: Networking'])
