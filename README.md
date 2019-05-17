# go-slice-utils

> A small package that contains common used functions for dealing with slices in go.

[![CircleCI](https://img.shields.io/circleci/project/brpaz/go-slice-utils.svg?style=for-the-badge)](https://circleci.com/gh/brpaz/go-slice-utils)
[![Code Climate](https://img.shields.io/codeclimate/maintainability/brpaz/go-slice-utils.svg?style=for-the-badge)](https://codeclimate.com/github/brpaz/go-slice-utils)
[![Code Climate coverage](https://img.shields.io/codeclimate/coverage/brpaz/go-slice-utils.svg?style=for-the-badge)](https://codeclimate.com/github/brpaz/go-slice-utils/progress/coverage)
![LICENSE](https://img.shields.io/github/license/brpaz/go-slice-utils.svg?style=for-the-badge)

## Functions

* Contains - Checks if a slice contains a particular element. Specific types variants are supported. Ex: ContainsString, ContainsInt etc

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

See [here](https://godoc.org/github.com/brpaz/go-slice-utils) for package documentation.

## Development


### Getting started

Make sure you have go installed and correctly configured on your machine.

This project uses [pre-commit](https://pre-commit.com/) to enforce coding style guidelines. Make sure you have it installed on your project:

```sh
pip install pre-commit
pre-commit install
```


### Run tests

```make tests```


## Contributing

- Fork it!
- Create your feature branch: git checkout -b my-new-feature
- Commit your changes: git commit -am 'Add some feature'
- Push to the branch: git push origin my-new-feature
- Submit a pull request :D

## License

[LICENSE](LICENSE)
