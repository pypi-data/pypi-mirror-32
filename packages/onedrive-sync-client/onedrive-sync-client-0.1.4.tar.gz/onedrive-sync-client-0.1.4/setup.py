# Copyright (C) 2018  XU Guang-zhao
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, only version 3 of the License, but not any
# later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

from pathlib import Path

from setuptools import setup, find_packages

setup(
    name='onedrive-sync-client',
    version='0.1.4',
    packages=find_packages(),
    install_requires=['requests-oauthlib'],
    author='XU Guang-zhao',
    description='OneDrive Client with Two-way Synchronizing Feature',
    license='AGPL-3.0-only',
    keywords='onedrive sync',
    url='https://github.com/gzxu/onedrive-sync-client',
    project_urls={
        'Bug Tracker': 'https://github.com/gzxu/onedrive-sync-client/issues',
        'Source Code': 'https://github.com/gzxu/onedrive-sync-client'
    },
    python_requires='>=3.5',
    entry_points={
        'console_scripts': [
            'onedrive = onedrive.main:main'
        ]
    },
    long_description=Path('README.md').read_text(),
    long_description_content_type='text/markdown',
    classifiers=(
        'Development Status :: 3 - Alpha',
        'Environment :: Console',
        'Intended Audience :: End Users/Desktop',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: GNU Affero General Public License v3',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
        'Programming Language :: Python :: 3.7',
        'Topic :: Office/Business :: Office Suites',
        'Topic :: Utilities'
    )
)
