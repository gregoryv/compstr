// Package strmix provides split functions for compounded words.
//
// A compounded word is made of two or more words merged together.
// Compound format is described at [Go initialisms].
//
// [Go initialisms]: https://github.com/golang/go/wiki/CodeReviewComments#initialisms
package strmix

import (
	"unicode"
	"unicode/utf8"
)

// USplit returns all non empty words from the given unicode input.
func USplit(s string) []string {
	dst := make([]string, 0, len(s)/2)
	return UAppendWords(dst, s)
}

// UAppendWords appends all non empty words from the given unicode input
// to dst. See [USplit] for how a string is divided into words.
func UAppendWords(dst []string, s string) []string {
	var from int
	var last rune
	var count int // number of letters since last split
	for i, r := range s {
		switch {
		case isUSeparator(r):
			if i > from {
				dst = append(dst, s[from:i])
				count = 0
			}
			from = i + 1

		case unicode.IsLower(last) && unicode.IsUpper(r):
			dst = append(dst, s[from:i])
			count = 0
			from = i

		case count >= 2 && unicode.IsUpper(last) && unicode.IsLower(r):
			to := i - utf8.RuneLen(last)
			dst = append(dst, s[from:to])
			count = 0
			from = to

		default:
			count++
		}
		last = r
	}
	return append(dst, s[from:])
}

func isUSeparator(r rune) bool {
	return r == '_' || r == ' ' || r == '.'
}
