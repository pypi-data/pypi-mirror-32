# fstring
[![Build Status](https://travis-ci.org/rinslow/fstring.svg?branch=master)](https://travis-ci.org/rinslow/fstring)
![PyPI - Python Version](https://img.shields.io/pypi/pyversions/fstring.svg)
[![PyPI version](https://badge.fury.io/py/fstring.svg)](https://badge.fury.io/py/fstring)

[PEP498](https://www.python.org/dev/peps/pep-0498/) Backport for all supported Python versions.

## Usage
```pip install fstring```

```python
from fstring import fstring as f
x = 1
y = 2
print f("{x}+{y}={x+y}")  # 1+2=3
```
