# Copyright (C) 2016 Duncan Macleod
#
# This program is free software; you can redistribute it and/or modify it
# under the terms of the GNU General Public License as published by the
# Free Software Foundation; either version 3 of the License, or (at your
# option) any later version.
#
# This program is distributed in the hope that it will be useful, but
# WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General
# Public License for more details.
#
# You should have received a copy of the GNU General Public License along
# with this program; if not, write to the Free Software Foundation, Inc.,
# 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.

"""setup for segments
"""

import os.path
import re
import sys

from setuptools import (setup, find_packages, Extension)


# get version
def find_version(path):
    with open(path, 'r') as fp:
        version_file = fp.read()
    version_match = re.search(r"^__version__ = ['\"]([^'\"]*)['\"]",
                              version_file, re.M)
    if version_match:
        return version_match.group(1)
    raise RuntimeError("Unable to find version string.")


# collect package data
packages = find_packages()

# declare dependencies
setup_requires = ['setuptools']
install_requires = ['six']

# add test dependencies
if set(('pytest', 'test', 'prt')).intersection(sys.argv):
    setup_requires.append('pytest_runner')

# define extension
csegments = Extension(
    'ligo.__segments',
    ['src/segments.c', 'src/infinity.c', 'src/segment.c', 'src/segmentlist.c'],
    include_dirs=['src'],
)

# run setup
setup(name='ligo-segments',
      version=find_version(os.path.join('ligo', 'segments.py')),
      description='Representations of semi-open intervals',
      author='Kipp Cannon',
      author_email='kipp.cannon@ligo.org',
      packages=packages,
      namespace_packages=['ligo'],
      setup_requires=setup_requires,
      install_requires=install_requires,
      ext_modules=[csegments],
      classifiers=[
          'Development Status :: 5 - Production/Stable',
          'Programming Language :: Python',
          'Programming Language :: Python :: 2.7',
          'Programming Language :: Python :: 3.4',
          'Programming Language :: Python :: 3.5',
          'Programming Language :: Python :: 3.6',
          'Intended Audience :: Science/Research',
          'Intended Audience :: End Users/Desktop',
          'Intended Audience :: Developers',
          'Natural Language :: English',
          'Topic :: Scientific/Engineering',
          'Topic :: Scientific/Engineering :: Astronomy',
          'Topic :: Scientific/Engineering :: Physics',
          'Operating System :: POSIX',
          'Operating System :: Unix',
          'Operating System :: MacOS',
          'License :: OSI Approved :: GNU General Public License v3 (GPLv3)',
      ],
      )
