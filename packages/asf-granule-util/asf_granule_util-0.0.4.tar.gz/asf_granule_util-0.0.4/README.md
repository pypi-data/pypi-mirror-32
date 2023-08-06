# asf_granule_util

## Docs

Basic documentation can be found [here](http://asf-granule-util-docs.s3-website-us-west-2.amazonaws.com/)
as well as example scripts which can be found in the repository under the `examples` directory.


## Install
Requirements are installed using [pipenv](https://docs.pipenv.org/):

    pipenv install

## Tests
To run tests

    ./run_all_tests.sh

Or run a particular test with

    pipenv run python tests/test_<test name>.py

## Pypi

To upload to pypi test

    ./upload.sh test


To upload to pypi

    ./upload.sh prod
