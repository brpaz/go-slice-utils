# go-slice-utils

> A small package that contains common used functions for dealing with slices in go.

[![CircleCI](https://img.shields.io/circleci/project/brpaz/go-slice-utils.svg?style=for-the-badge)](https://circleci.com/gh/brpaz/go-slice-utils)
[![Code Climate](https://img.shields.io/codeclimate/maintainability/brpaz/go-slice-utils.svg?style=for-the-badge)](https://codeclimate.com/github/brpaz/go-slice-utils)
[![Code Climate coverage](https://img.shields.io/codeclimate/coverage/brpaz/go-slice-utils.svg?style=for-the-badge)](https://codeclimate.com/github/brpaz/go-slice-utils/progress/coverage)
![LICENSE](https://img.shields.io/github/license/brpaz/go-slice-utils.svg?style=for-the-badge)

**This is work in progress. use it at your own risk**

## Functions

* Contains - Checks if a slice contains a particular element.
* ContainsBool
* ContainsString
* ContainsInt

...

## Install

```
go get -u github.com/brpaz/go-slice-utils/slice
```

## Usage

```go
import "github.com/brpaz/go-slice-utils/slice"
slice := []string{"one", "two", "three"}
slice.ContainsString(slice, "two")
```

## Development

### Run tests

```make tests```

## License

[LICENSE](LICENSE)
