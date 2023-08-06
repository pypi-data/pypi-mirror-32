import os
from setuptools import setup, find_packages

install_requires = [line.rstrip() for line in open(os.path.join(os.path.dirname(__file__), "requirements.txt"))]

setup(name='proforma',
      version='0.2.1',
      description='Simple library for setting up AWS Infrastructure.',
      url='http://github.com/HumanCellAtlas/proforma',
      author='Sam Pierson',
      author_email='spierson@chanzuckerberg.com',
      license='MIT',
      packages=find_packages(exclude=['tests']),
      zip_safe=False,
      install_requires=install_requires,
      platforms=['MacOS X', 'Posix'],
      test_suite='test',
      classifiers=[
          'Intended Audience :: Developers',
          'License :: OSI Approved :: MIT License',
          'Operating System :: MacOS :: MacOS X',
          'Operating System :: POSIX',
          'Programming Language :: Python :: 3.6'
      ]
      )
