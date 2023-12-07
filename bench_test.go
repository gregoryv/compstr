package strmix

import (
	"testing"

	"github.com/fatih/camelcase"
)

const in = "CallRPCHandler"

func BenchmarkSplit_strmix(b *testing.B)    { bench(b, Split) }
func BenchmarkSplit_camelcase(b *testing.B) { bench(b, camelcase.Split) }
func BenchmarkUSplit_strmix(b *testing.B)   { bench(b, USplit) }

func bench(b *testing.B, split func(string) []string) {
	for i := 0; i < b.N; i++ {
		_ = split(in)
	}
}

func BenchmarkAppendWords_strmix(b *testing.B) {
	in := "CallRPCHandler"
	strmix := make([]string, 0, 3)
	for i := 0; i < b.N; i++ {
		_ = AppendWords(strmix, in)
	}
}
