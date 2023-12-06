package compstr

import (
	"testing"
	// "github.com/fatih/camelcase"
)

const in = "CallRPCHandler"

func Benchmark_compstr(b *testing.B) { bench(b, Split) }

//func Benchmark_camelcase(b *testing.B)    { bench(b, camelcase.Split) }

func bench(b *testing.B, split func(string) []string) {
	for i := 0; i < b.N; i++ {
		_ = split(in)
	}
}

func BenchmarkAppendWords(b *testing.B) {
	in := "CallRPCHandler"
	compstr := make([]string, 0, 3)
	for i := 0; i < b.N; i++ {
		_ = AppendWords(compstr, in)
	}
}
