qtumsuite
=========

[![Build Status](http://img.shields.io/travis/btcsuite/qtumsuite.svg)](https://travis-ci.org/btcsuite/qtumsuite)
[![Coverage Status](http://img.shields.io/coveralls/btcsuite/qtumsuite.svg)](https://coveralls.io/r/btcsuite/qtumsuite?branch=master)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/qtumproject/qtumsuite)

package qtumsuite provides bitcoin-specific convenience functions and types.
A comprehensive suite of tests is provided to ensure proper functionality.  See
`test_coverage.txt` for the gocov coverage report.  Alternatively, if you are
running a POSIX OS, you can run the `cov_report.sh` script for a real-time
report.


## Installation and Updating

```bash
$ go get -u github.com/qtumproject/qtumsuite
```

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from the btcsuite developers.  To
verify the signature perform the following:

- Download the public key from the Conformal website at
  https://opensource.conformal.com/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

package qtumsuite is licensed under the [copyfree](http://copyfree.org) ISC
License.
