# armorvise
[![PyPI version](https://img.shields.io/pypi/v/armorvise.svg)](https://pypi.python.org/pypi/armorvise/)

[armorvise](https://github.com/numberwolf/armorvise) is a daemon tool on linux.

```console
$ armorvise "bash script.sh >/dev/null 2>&1"
$ armorvise -p
All Proc
PID CMD
12048 /usr/bin/python /usr/local/bin/armorvise bash script.sh >/dev/null 2>&1 
```

## Installation

### Prerequisites

* **[Python 2](https://www.python.org/downloads/)**
* **[Linux](https://www.linux.org/)**

### Option 1: Install via pip

    $ pip install armorvise

### Option 2: Download from GitHub

```
$ [sudo] git clone https://github.com/numberwolf/armorvise.git
$ cd armorvise
$ python setup.py install
```

## Upgrading

```
$ pip install --upgrade armorvise
```






