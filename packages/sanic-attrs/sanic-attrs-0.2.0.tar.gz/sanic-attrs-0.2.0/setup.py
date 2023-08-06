import os
from setuptools import setup

here = os.path.abspath(os.path.dirname(__file__))

about = {'version': '0.2.0'}

with open(
    os.path.join(here, 'sanic_attrs', '__init__.py'), 'r', encoding='utf-8'
) as f:
    for line in f:
        if line.startswith('__version__'):
            about['version'] = line.strip().split('=')[1].strip(' \'"')

with open(os.path.join(here, 'README.md'), 'r', encoding='utf-8') as f:
    README = f.read()


setup(
    name='sanic-attrs',
    version=about['version'],
    url='http://github.com/vltr/sanic-attrs/',
    license='MIT',
    author='Channel Cat',
    author_email='channelcat@gmail.com',
    maintainer='Richard Kuesters',
    maintainer_email="rkuesters@gmail.com",
    description='OpenAPI / Swagger support for Sanic using attrs',
    long_description=README,
    long_description_content_type='text/markdown',
    packages=['sanic_attrs'],
    package_data={'sanic_attrs': ['ui/*']},
    platforms='any',
    install_requires=[
        'sanic>=0.7.0',
        'attrs>=18.0.0',
    ],
    classifiers=[
        'Development Status :: 2 - Pre-Alpha',
        'Environment :: Web Environment',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
        'Programming Language :: Python :: 3.7',
    ],
)
