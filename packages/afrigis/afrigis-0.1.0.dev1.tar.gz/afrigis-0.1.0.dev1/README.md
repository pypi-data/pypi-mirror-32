# Afrigis Python Library

[![Maintainability](https://api.codeclimate.com/v1/badges/ecd8973da991fa593b82/maintainability)](https://codeclimate.com/github/chris-cmsoft/python-afrigis/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/ecd8973da991fa593b82/test_coverage)](https://codeclimate.com/github/chris-cmsoft/python-afrigis/test_coverage)
[![Build Status](https://travis-ci.org/chris-cmsoft/python-afrigis.svg?branch=master)](https://travis-ci.org/chris-cmsoft/python-afrigis)

### Installation

```bash
$ pip install afrigis
```

### Services:

#### Geocode

Example on using the Geocode service
```python
from afrigis.services import geocode
result = geocode('AFRIGIS_KEY', 'AFRIGIS_SECRET', 'NADID | SEOID')
print(result)
# {'number_of_records': 4, ...}
```

### Running tests

```bash
$ py.test
```

### Building and pushing to pypi

> In order to di this please make sure you are authenticated against Pypi first :).  
> You ca do this with the following method: https://docs.python.org/3/distutils/packageindex.html#the-pypirc-file

```bash
$ python setup.py sdist upload
```