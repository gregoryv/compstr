compstr - Provides means to split compound words

## Benchmark

    goos: linux
    goarch: amd64
    pkg: compstr
    cpu: Intel(R) Xeon(R) E-2288G CPU @ 3.70GHz
    Benchmark_compstr
    Benchmark_compstr-16              16188284       140.7  ns/op     112 B/op       1 allocs/op
    Benchmark_faithCamelcase-16        1367502      1007    ns/op     488 B/op      24 allocs/op
    Benchmark_proprotoCamelcase-16     1448004       776.5  ns/op     488 B/op      24 allocs/op
    BenchmarkAppendWords-16           36377338        27.94 ns/op       0 B/op       0 allocs/op


## Related modules

- [faith/camelcase](https://pkg.go.dev/github.com/fatih/camelcase)
- [proproto/camelcase](https://pkg.go.dev/github.com/proproto/camelcase)
