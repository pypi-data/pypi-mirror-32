# [Sematext API](http://sematext-api.hive.pt)

Python API client for Sematext.

## Configuration

* `SEMATEXT_TOKEN` (`str`) - The secret token to be for authentication on the Sematext API (defaults to `None`)
* `SEMATEXT_BUFFER_SIZE` (`int`) - The size of the buffer (in number of entries) until the buffer is flushed (defaults to `128`)
* `SEMATEXT_TIMEOUT` (`int`) - The timeout in seconds in seconds until the buffer is flushed (defaults to `30`)

## License

Sematext API is currently licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/).

## Build Automation

[![Build Status](https://travis-ci.org/hivesolutions/sematext_api.svg?branch=master)](https://travis-ci.org/hivesolutions/sematext_api)
[![Coverage Status](https://coveralls.io/repos/hivesolutions/sematext_api/badge.svg?branch=master)](https://coveralls.io/r/hivesolutions/sematext_api?branch=master)
[![PyPi Status](https://img.shields.io/pypi/v/sematext_api.svg)](https://pypi.python.org/pypi/sematext_api)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://www.apache.org/licenses/)
