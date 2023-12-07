package strmix

import (
	"fmt"
	"slices"
	"testing"
)

func Example_unicodeSplit() {
	fmt.Println(USplit("ÖstraÅkern"))
	fmt.Println(USplit("1976_årskurs"))
	fmt.Println(USplit("lilla.örnenFlög"))
	fmt.Println(USplit("XYZÖken"))
	// output:
	// [Östra Åkern]
	// [1976 årskurs]
	// [lilla örnen Flög]
	// [XYZ Öken]
}

func TestUSplit(t *testing.T) {
	cases := []struct {
		in  string
		exp []string
	}{
		{in: "", exp: []string{""}},
		{in: "lowercase", exp: []string{"lowercase"}},
		{in: "Class", exp: []string{"Class"}},
		{in: "MyClass", exp: []string{"My", "Class"}},
		{in: "MyC", exp: []string{"My", "C"}},
		{in: "HTML", exp: []string{"HTML"}},
		{in: "TestSplit_format", exp: []string{"Test", "Split", "format"}},
		{in: "PDFLoader", exp: []string{"PDF", "Loader"}},
		{in: "AString", exp: []string{"A", "String"}},
		{in: "SimpleXMLParser", exp: []string{"Simple", "XML", "Parser"}},
		{in: "vimRPCPlugin", exp: []string{"vim", "RPC", "Plugin"}},
		{in: "GL11Version", exp: []string{"GL11", "Version"}},
		{in: "99Bottles", exp: []string{"99", "Bottles"}},
		{in: "May5", exp: []string{"May5"}},
		{in: "BFG9000", exp: []string{"BFG9000"}},
		{in: "BöseÜberraschung", exp: []string{"BöseÜberraschung"}},
		{in: "Two  spaces", exp: []string{"Two", "spaces"}},
		{in: "http.ResponseWriter", exp: []string{"http", "Response", "Writer"}},
		{in: "BadUTF8\xe2\xe2\xa1", exp: []string{"Bad", "UTF8\xe2\xe2\xa1"}},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			got := Split(c.in)
			if !slices.Equal(got, c.exp) {
				t.Error("got:", got, "exp:", c.exp)
			}
		})
	}
}
