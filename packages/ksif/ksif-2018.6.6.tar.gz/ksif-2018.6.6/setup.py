# -*- coding: utf-8 -*-
"""
:Author: Jaekyoung Kim
:Date: 2018. 6. 6.
"""
from setuptools import setup, find_packages

with open("README.md", "r") as fh:
    long_description = fh.read()

setup(
    name             = 'ksif',
    version          = '2018.6.6',
    description      = 'Quantitative investment tools for KSIF',
    long_description = long_description,
    author           = 'Jaekyoung Kim',
    author_email     = 'jaekyoungkim@kaist.ac.kr',
    url              = 'https://github.com/willbelucky/ksif',
    download_url     = 'https://github.com/willbelucky/ksif/archive/master.zip',
    install_requires = ['pandas>=0.21.1', 'tqdm>=4.19.5'],
    packages         = find_packages(exclude = ['tests*']),
    keywords         = ['ksif', 'portfolio', 'backtest', 'finance'],
    python_requires  = '>=3',
    classifiers=(
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ),
)