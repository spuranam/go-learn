// Attributions
// some of the details below have been reproduced here from;
// Effective Go [http://golang.org/doc/effective_go.html]
// The Go Programming Language Specification [http://golang.org/ref/spec]

package main

import "fmt"

func main() {
	// defers can stacked, each defer would be called when the function exits
	// in LIFO (last-in-first-out) order
	for i := 0; i < 10; i++ {
		defer fmt.Printf("Calling defered statement %d\n", i)
	}
	fmt.Printf("The quick brown fox\n")
}
