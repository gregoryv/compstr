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
    BenchmarkSplit_strmix-16          15030070    119.70 ns/op     112 B/op       1 allocs/op
    BenchmarkSplit_camelcase-16        1237680    937.60 ns/op     488 B/op      24 allocs/op
    BenchmarkAppendWords_strmix-16    48409911     21.00 ns/op       0 B/op       0 allocs/op


## Related modules

- [faith/camelcase](https://pkg.go.dev/github.com/fatih/camelcase)

