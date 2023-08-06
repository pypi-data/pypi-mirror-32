import setuptools
from EnumSubDomain import __version__

with open('README.md', 'r') as f:
    long_description = f.read()

setuptools.setup(
    name='ESD',
    version=__version__,
    author='Feei',
    author_email='feei@feei.cn',
    description='Enumeration sub domains(枚举子域名)',
    long_description=long_description,
    long_description_content_type='text/markdown',
    url='https://github.com/FeeiCN/ESD',
    packages=setuptools.find_packages(),
    install_requires=[
        'aiodns',
        'aiohttp',
        'async-timeout',
        'colorlog',
        'requests'
    ],
    classifiers=(
        "Topic :: Security",
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: GNU General Public License v3 (GPLv3)"
    ),
    include_package_data=True,
    package_data={
        '': ['*.esd']
    },
    entry_points={
        'console_scripts': [
            'esd=EnumSubDomain:main'
        ]
    }
)
