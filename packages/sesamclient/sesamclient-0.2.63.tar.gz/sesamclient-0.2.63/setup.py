# Copyright (C) Bouvet ASA - All Rights Reserved.
# Unauthorized copying of this file, via any medium is strictly prohibited.
import os

try:  # for pip >= 10
    from pip._internal.req import parse_requirements
    from pip._internal.download import PipSession
except ImportError: # for pip <= 9.0.3
    from pip.req import parse_requirements
    from pip.download import PipSession


from setuptools import setup, find_packages

install_reqs = parse_requirements("requirements.txt", session=PipSession())
requires = [str(ir.req) for ir in install_reqs if ir.req is not None]

here = os.path.abspath(os.path.dirname(__file__))

with open(os.path.join(here, 'README.txt')) as f:
    README = f.read()

with open(os.path.join(here, "sesamclient", 'VERSION.txt')) as f:
    VERSION = f.read().strip()

setup(name='sesamclient',
      version=VERSION,
      description='sesamapi client',
      long_description=README,
      classifiers=[
          "Programming Language :: Python",
          "Topic :: Internet :: WWW/HTTP",
      ],
      author='Sesam',
      url='http://sesam.io',
      packages=find_packages(),
      include_package_data=True,
      zip_safe=False,
      setup_requires=['pip'],
      test_suite="nose.collector",
      install_requires=requires,
      entry_points={
          'console_scripts': [
               'sesam=sesamclient.main.main:main',
          ],
      })
