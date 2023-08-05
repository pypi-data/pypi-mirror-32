from setuptools import setup
#from distutils.core import setup
import os
import subprocess

#with open('fastent/requirements.txt') as f:
 #   requirements = f.read().splitlines()

#print(requirements)
def readme():
    with open('README.rst') as f:
        return f.read()

setup(
  name = 'fastent',
  packages = ['fastent'], # this must be the same as the name above
  version = '0.5.8.6',
  description = 'Automated Custom NER tool',
  author = 'Erik Arakelyan',
  author_email = 'erikarakelyan1997@gmail.com',
  scripts=['fastent/fastent_install.sh'],
  install_requires=['alabaster==0.7.10', 'argh==0.26.2', 'Babel==2.5.3', 'boto==2.48.0', 'boto3==1.7.24', 'botocore==1.10.24', 'bz2file==0.98', 'certifi==2018.4.16', 'chardet==3.0.4', 'CouchDB==1.2', 'cymem==1.31.2', 'cytoolz==0.8.2', 'dill==0.2.7.1', 'docutils==0.14', 'fuzzywuzzy==0.16.0', 'gensim==3.4.0', 'idna==2.6', 'imagesize==1.0.0', 'Jinja2==2.10', 'jmespath==0.9.3', 'livereload==2.5.2', 'lxml==4.2.1', 'MarkupSafe==1.0', 'msgpack-numpy==0.4.1', 'msgpack-python==0.5.6', 'murmurhash==0.28.0', 'nltk==3.3', 'numpy==1.14.3', 'packaging==17.1', 'pandas==0.23.0', 'pathlib==1.0.1', 'pathtools==0.1.2', 'plac==0.9.6', 'port-for==0.3.1', 'praw==5.4.0', 'prawcore==0.14.0', 'preshed==1.0.0', 'Pygments==2.2.0', 'pyparsing==2.2.0', 'python-dateutil==2.7.3', 'python-Levenshtein==0.12.0', 'pytz==2018.4', 'PyYAML==3.12', 'regex==2017.4.5', 'requests==2.18.4', 's3transfer==0.1.13', 'scipy==1.1.0', 'six==1.11.0', 'smart-open==1.5.7', 'snowballstemmer==1.2.1', 'spacy==2.0.11', 'Sphinx==1.7.4', 'sphinx-autobuild==0.7.1', 'sphinxcontrib-websupport==1.0.1', 'termcolor==1.1.0', 'thinc==6.10.2', 'toolz==0.9.0', 'tornado==5.0.2', 'tqdm==4.23.3', 'ujson==1.35', 'update-checker==0.16', 'urllib3==1.22', 'watchdog==0.8.3', 'wrapt==1.10.11']
,
  url = 'https://github.com/fastent/fastent', # use the URL to the github repo
  long_description=readme(),
  download_url = 'https://github.com/fastent/fastent/archive/0.5.tar.gz', # I'll explain this in a second
  keywords = ['NER', 'Anotation', 'Contextualization'], # arbitrary keywords
  classifiers = [],
)


