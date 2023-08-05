# encoding: utf-8
from setuptools import setup, find_packages
import sys, os

VERSION = '0.1.0'

setup(name='pyfi_helper',
      version=VERSION,
      description='encapsulate the windPy API with pandas',
      url='https://github.com/wangluzhou/pyfi.git',
      author='wangluzhou',
      author_email='wangluzhou@aliyun.com',
      license='MIT',
      packages=find_packages(),
      zip_safe=True)
