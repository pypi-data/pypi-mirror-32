 [![Build Status](https://travis-ci.org/SimplySecurity/SimplyDomain.svg?branch=master)](https://travis-ci.org/SimplySecurity/SimplyDomain)
  [![Coverage Status](https://coveralls.io/repos/github/SimplySecurity/SimplyDomain/badge.svg?branch=master)](https://coveralls.io/github/SimplySecurity/SimplyDomain?branch=master)
 ![Alt text](simplydomain/docs/SimplyDomain-logo.png?raw=true "SimplyDomain")

 
# simplydomain-pkg 
Subdomain brute force focused on speed and data serialization. 
SimplyDomain uses a framework approach to build and deploy modules within. This allows
for fast, easy and concise output to feed into larger OSINT feeds.

## Three Core Design Principals:
* Easy install - support as many *NIX* bassed platforms.
* Pure Python - no other arbitrary set up processes and Python-3 support
* Expose public API - allows for simplydomain to intergrate into other toolsets.

# simplydomain Install

## Using PIP Package Managment
```python
pip3 install simplydomain
```
or 
```python
python3 -m pip install simplydomain
```
## Using `setup.py` to build from source
```bash
git clone git@github.com:SimplySecurity/simplydomain-pkg.git | cd simplydomain-pkg
python3 -m pip install .
```
## simplydomain CLI tools (Quickstart)
simplydomain supports a `bin` directory which are installed during the Python Setup PKG install. This now allows users to simply use their terminal of choice to use simplydomain.

To display Help:
```bash
simply_domain.py -h 
```
To run a basic passive sub-domain search:
```bash
simply_domain.py -all uber.com
```
## simplydomain programming API 
simplydomain now supports a full 

