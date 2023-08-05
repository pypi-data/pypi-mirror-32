from setuptools import setup, find_packages
from codecs import open
from os import path

__version__ = '0.1.3'

here = path.abspath(path.dirname(__file__))

# Get the long description from the README file
with open(path.join(here, 'readme.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
    name='10daysweb',
    version=__version__,
    description='Async web framework for learning',
    long_description=long_description,
    url='https://github.com/bace1996/10daysWeb',
    author='Cykrt Chan',
    author_mail='cykrt1996@gmail.com',
    license='MIT',
    classifiers=[
        'Development Status :: 3 - Alpha',
        'Intended Audience :: Developers',
        'Environment :: Web Environment',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3.6',
    ],
    keywords='web framework async',
    packages=find_packages(exclude=['docs', 'demos', 'tests*']),
    include_package_data=True,
    install_requires=[
        'httptools',
    ],
    python_requires='>=3.6',
)
