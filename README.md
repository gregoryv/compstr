compstr - Provides means to split compound words

## Benchmark

    goo: linux
    goarch: amd64
    pkg: compstr
    cpu: Intel(R) Xeon(R) E-2288G CPU @ 3.70GHz
    BenchmarkSplit-16           9570945   128.20 n/op   112 B/op    1 alloc/op
    BenchmarkAppendWord-16     42347318    28.10 n/op     0 B/op    0 alloc/op
