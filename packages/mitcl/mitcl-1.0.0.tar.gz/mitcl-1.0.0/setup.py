"""
set up tools for mitcl
"""
from setuptools import setup, find_packages
from codecs import open
from os import path

here = path.abspath(path.dirname(__file__))

setup(
    #
    include_package_data=True,
    name='mitcl',
    version='1.0.0',
    description='A tool for MIT OCW check and download',
    url='https://github.com/alipython/mitcl',
    author='noxaean',
    author_email='noxaean@gmail.com',
    license='MIT',
    classifiers=[
        'Development Status :: 3 - Alpha',
        'Topic :: Software Development :: Build Tools',
    	'Programming Language :: Python :: 2',
    	'Programming Language :: Python :: 2.7',
    	'Programming Language :: Python :: 3',
    	'Programming Language :: Python :: 3.2',
    	'Programming Language :: Python :: 3.3',
    	'Programming Language :: Python :: 3.4',
        ],
    keywords='mit ocw check',
    packages=find_packages('mitcl'),
    package_dir={'':'mitcl'},
    package_data={
	'': ['*.rst', '*.dat', '*.txt']
       },
    py_modules=["mitcl"],
    entry_points={
        'console_scripts':[
            'mitcl=mitcl:main',
            ],
        }
    )
