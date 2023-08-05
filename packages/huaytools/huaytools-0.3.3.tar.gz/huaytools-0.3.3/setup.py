"""
* Project Name: huaytools
* Author: huay
* Mail: imhuay@163.com
* Created Time:  2018-1-26 11:33:13
"""
from setuptools import setup, find_packages
from huaytools import __version__

version = __version__

install_requires = [
    'six',
    'bs4',
    # 'bunch',  # add the hole bunch in, and apply it to Python3

    # Use Anaconda to install avoid install failed
    # 'tensorflow',  # 1.6
    # 'gensim',
    # 'numpy',
]

package_data = {'huaytools.nlp': ['data/*']}

setup(
    name="huaytools",
    version=version,
    keywords=("huay", "huaytools"),
    description="huay's tools",
    long_description="huay's tools",
    license="MIT Licence",
    url="https://github.com/imhuay/huaytools",
    author="huay",
    author_email="imhuay@163.com",
    packages=find_packages(),
    package_data=package_data,
    # include_package_data=True,
    platforms="any",
    install_requires=install_requires
)
