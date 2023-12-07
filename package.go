// Package strmix provides split functions for compounded words.
//
// A compounded word is made of two or more words merged together.
// Compound format is described at [Go initialisms].
//
// [Go initialisms]: https://github.com/golang/go/wiki/CodeReviewComments#initialisms
package strmix

// Split returns all non empty words from the given ASCII input.
func Split(s string) []string {
	dst := make([]string, 0, len(s)/2)
	return AppendWords(dst, s)
}

// AppendWords appends all non empty words from the given ASCII input
// to dst. See [Split] for how a string is divided into words.
func AppendWords(dst []string, s string) []string {
	var from int
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '_':
			if i > from {
				dst = append(dst, s[from:i])
			}
			from = i + 1

		case isLower(s[i-1]) && isUpper(s[i]):
			dst = append(dst, s[from:i])
			from = i

		case i-from > 2 && isUpper(s[i-1]) && isLower(s[i]):
			dst = append(dst, s[from:i-1])
			from = i - 1
		}
	}
	return append(dst, s[from:])
}

func isLower(b byte) bool { return b >= 'a' && b <= 'z' }
func isUpper(b byte) bool { return b >= 'A' && b <= 'Z' }
