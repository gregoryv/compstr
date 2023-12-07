package strmix

import (
	"fmt"
	"slices"
	"testing"
)

func Example_format() {
	fmt.Println(Split("ServeHTTP"))
	fmt.Println(Split("CallRPCMethod"))
	fmt.Println(Split("ExampleCar_Model"))
	// output:
	// [Serve HTTP]
	// [Call RPC Method]
	// [Example Car Model]
}

func ExampleSplit() {
	fmt.Println(Split("CallRPCHandler_nowOrLater"))
	fmt.Println(Split("H20Chemical"))
	fmt.Println(Split("hiddenField"))
	fmt.Println(Split("word"))
	fmt.Println(Split("ab__cd"))
	fmt.Println(Split("1_2_3A"))
	fmt.Println(Split("http.ResponseWriter"))
	// output:
	// [Call RPC Handler now Or Later]
	// [H20 Chemical]
	// [hidden Field]
	// [word]
	// [ab cd]
	// [1 2 3A]
	// [http Response Writer]
}

func TestSplit(t *testing.T) {
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
