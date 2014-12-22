// Attributions
// some of the details below have been reproduced here from;
// Effective Go [http://golang.org/doc/effective_go.html]
// The Go Programming Language Specification [http://golang.org/ref/spec]

package main

import "fmt"

// Type switches are a form of conversion: they take an interface and,
// for each case in the switch, in a sense convert it to the type of
// that case. Here's a simplified version of how the code under
// fmt.Printf turns a value into a string using a type switch.
// If it's already a string, we want the actual string value held by
// the interface, while if it has a String method we want the result
// of calling the method.

type Stringer interface {
	String() string
}

type mockString struct {
	content string
}

func (s *mockString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

func main() {
	s := &mockString{"The quick brown fox jumps over the lazy dog"}
	str := "Once upon a time in a land far far away"
	printString(s)
	printString(str)
}
