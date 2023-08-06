# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from setuptools import find_packages, setup

setup(
    name='dadata-client',
    version='0.1.0',
    description="Dadata.ru python client",
    keywords=[],
    url="https://github.com/f213/dadata-client/",
    author="Fedor Borshev",
    author_email="f@f213.in",
    license="MIT",
    packages=find_packages(),
    install_requires=[
        'requests',
    ],
    include_package_data=True,
    zip_safe=False,
)
