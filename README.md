# `gdt-examples` - The Go Declarative Testing framework - Fixtures library ![go test workflow](https://github.com/jaypipes/gdt-examples/actions/workflows/gate-tests.yml/badge.svg)

[`gdt`][gdt] is a testing library that allows test authors to cleanly describe tests
in a YAML file. `gdt` reads YAML files that describe a test's assertions and
then builds a set of Go structures that the standard Go
[`testing`](https://golang.org/pkg/testing/) package can execute.

[gdt]: https://github.com/jaypipes/gdt

`gdt-examples` is a companion Go library for `gdt` that contains examples for
how to use `gdt` and its companion Go libraries like `gdt-http` and
`gdt-fixtures`.

## Installation

`gdt-examples` is a Go library and is intended to be included in your own Go
application's test code as a Go package dependency.

Install `gdt-examples` into your `$GOPATH` by executing:

```
go get -u github.com/jaypipes/gdt-examples
```

Alternately, include `github.com/jaypipes/gdt-examples` in your Go dependency
management of choice.

## Contributing and acknowledgements

`gdt` was inspired by [Gabbi](https://github.com/cdent/gabbi), the excellent
Python declarative testing framework. `gdt` tries to bring the same clear,
concise test definitions to the world of Go functional testing.

Contributions to `gdt-examples` are welcomed! Feel free to open a Github issue
or submit a pull request.
