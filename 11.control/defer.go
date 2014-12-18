// Attributions
// some of the details below have been reproduced here from;
// Effective Go [http://golang.org/doc/effective_go.html]
// The Go Programming Language Specification [http://golang.org/ref/spec]

package main

import "fmt"

func main() {
	// A "defer" statement invokes a function whose execution is deferred to
	// the moment the surrounding function returns, either because the
	// surrounding function executed a return statement, reached the end of
	// its function body, or because the corresponding goroutine is panicking.
	defer fmt.Printf(" jumps over the lazy dog\n")
	fmt.Printf("The quick brown fox")
}
