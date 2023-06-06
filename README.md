# Go Declarative Testing - Examples ![go test workflow](https://github.com/jaypipes/gdt-examples/actions/workflows/gate-tests.yml/badge.svg)

[`gdt`][gdt] is a testing library that allows test authors to cleanly describe tests
in a YAML file. `gdt` reads YAML files that describe a test's assertions and
then builds a set of Go structures that the standard Go
[`testing`](https://golang.org/pkg/testing/) package can execute.

[gdt]: https://github.com/jaypipes/gdt

This `gdt-examples` repository is a companion Go library for `gdt` that
contains examples for how to use `gdt` and its companion Go libraries like
`gdt-http` and `gdt-fixtures`.

The `http` directory in this repository contains Go code for a simplistic HTTP
web service that lists and creates books (in `http/api`) along with functional
tests written in YAML using `gdt`'s test file format (in `http/tests/api`).

## Contributing and acknowledgements

`gdt` was inspired by [Gabbi](https://github.com/cdent/gabbi), the excellent
Python declarative testing framework. `gdt` tries to bring the same clear,
concise test definitions to the world of Go functional testing.

Contributions to `gdt-examples` are welcomed! Feel free to open a Github issue
or submit a pull request.
