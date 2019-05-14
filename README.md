# go-slice-utils

> A small package that contains common used functions for dealing with slices in go.

[![CircleCI](https://img.shields.io/circleci/project/brpaz/gp-slice-utils.svg?style=for-the-badge)](https://circleci.com/gh/brpaz/go-slice-utils)
[[![Code Climate](https://img.shields.io/codeclimate/github/brpaz/go-slice-utils.svg?style=for-the-badge)](https://codeclimate.com/brpaz/go-slice-utils)](https://img.shields.io/codeclimate/maintainability/brpaz/go-slice-utils.svg?style=for-the-badge)
![LICENSE](https://img.shields.io/github/license/brpaz/go-slice-utils.svg?style=for-the-badge)

**This is work in progress. use it at your own risk**

## Functions

* Contains - Checks if a slice contains a particular element.
* ContainsBool
* ContainsString
* ContainsInt

...

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
