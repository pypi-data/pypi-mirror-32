# python-data-utils

[![CircleCI](https://circleci.com/gh/raywhite/python-data-utils.svg?style=shield&circle-token=97844f41e83cf8b384ee340dab4c201754467857)](https://circleci.com/gh/raywhite/python-data-utils)

> Common functions for typing and manipulating structured (but generally dirty) data.

## About

This is entirely a repository to hold any common functions and funcional utilities that might be used while cleaning and processing various bits of data coming form the many feeds we deal with @raywhite...

The code is just here to avoid duplication, and generally the docstrings are enough to figure out what any given function does (`__help__` is your friend). 

## Testing

The python environment is best set up using a virtual environment for handling pinning of dependancies, and the correct python version (`v2.7`). In order to create the correct virtual environment, you can do something like this (on mac, see the circle configuration for linux instructions);

```sh
$ virtualenv --python /usr/bin/python2.7 ./venv
$ chmod +x ./venv/bin/activate
$ . ./venv/bin/activate
$ pip install -r requirements.txt
$ pip install -r requirements-dev.txt
$ echo "do some stuff... like run tests"
```

Since it appears to be idiomatic in python - each test file is named for the corresponding file being tested, and is located in the same directory (`src`). To run all python tests use; `python -m unittest discover ./src "*test.py"`.

Lint can be checked using [pylint](https://www.pylint.org/); `pylint -r n ./src ./setup.py`.

## License

&bull; **MIT** &copy; Ray White, 2017-2018 &bull;
