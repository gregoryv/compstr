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
    BenchmarkSplit_strmix-16          17144353       109.6 ns/op     112 B/op       1 allocs/op
    BenchmarkSplit_camelcase-16        1138200       994.4 ns/op     488 B/op      24 allocs/op
    BenchmarkUSplit_strmix-16         10438047       153.3 ns/op     112 B/op       1 allocs/op
    BenchmarkAppendWords_strmix-16    58247101        20.5 ns/op       0 B/op       0 allocs/op


## Related modules

- [faith/camelcase](https://pkg.go.dev/github.com/fatih/camelcase)

