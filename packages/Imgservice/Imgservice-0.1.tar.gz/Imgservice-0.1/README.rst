Image service
#############

Library to create image proxying services.

The most basic usecase is serving images through HTTPS and stripping
information from the original request, similar to what
[camo](https://github.com/atmos/camo) does.

Imgservice brings it a step forward by allowing:

- Non-http(s) sources (eg. private S3 buckets, ...)
- Image processing (resizing, cropping, ...)


URL format and security
=======================

URLs are in the format: `https://example.com/<token>`, where `token`
is a (signed, b64 encoded) object describing the image source and any
required transformations.


Configuration
=============

All the configuration is done by writing a simple Python script that
imports and makes use of the imgservice library.


Testing
=======

To run the test suite:

    docker-compose run --rm -e PYTHONHASHSEED=0 web py.test -vvv --cov=imgservice --cov-report=term-missing ./tests
