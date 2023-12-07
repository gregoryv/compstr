[strmix](https://pkg.go.dev/github.com/gregoryv/strmix) - Provides means to split compound words

A compounded word is made of two or more words merged together.
Compound format is described at [Go initialisms](
https://github.com/golang/go/wiki/CodeReviewComments#initialisms
).

The primary use case is to efficiently split a string like
callRPCHandler into the words [call RPC Handler] during static Go
source code analysis.

This module is optimized for words comprised of characters
[a-zA-Z0-9_].

## Benchmark

    goos: linux
    goarch: amd64
    pkg: github.com/gregoryv/strmix
    cpu: Intel(R) Xeon(R) E-2288G CPU @ 3.70GHz
    BenchmarkSplit_strmix
    BenchmarkSplit_strmix-16       16483531       105.20 ns/op     112 B/op       1 allocs/op
    BenchmarkSplit_camelcase-16     1233808       967.80 ns/op     488 B/op      24 allocs/op
    BenchmarkUSplit_strmix-16       6354415       199.50 ns/op     112 B/op       1 allocs/op
    BenchmarkAppendWords-16        51060759        20.92 ns/op       0 B/op       0 allocs/op
    BenchmarkUAppendWords-16       22820806        52.30 ns/op       0 B/op       0 allocs/op

## Related modules

- [faith/camelcase](https://pkg.go.dev/github.com/fatih/camelcase)

