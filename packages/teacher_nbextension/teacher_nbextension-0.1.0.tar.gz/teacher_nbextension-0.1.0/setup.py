# from nbsetuptools import setup
from setuptools import setup
from os.path import abspath, dirname, join

NAME = 'teacher_nbextension'
AUTHOR = 'ibelyalov'
EMAIL = 'ibelyalov@yandex.ru'
DESCRIPTION = 'Teacher extension for Jupyter'
URL = 'http://github.com/theotheo/teacher_nbextension'

setup(
    name=NAME,
    description=DESCRIPTION,
    author=AUTHOR,
    author_email=EMAIL,
    url=URL,
    version="0.1.0",
    static=join(abspath(dirname(__file__)), 'static'),
    install_requires=[
        'python-logstash'
    ],
    include_package_data=True
)
