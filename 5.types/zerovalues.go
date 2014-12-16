// Attributions
// some of the details below have been reproduced here from;
// Effective Go [http://golang.org/doc/effective_go.html]
// The Go Programming Language Specification [http://golang.org/ref/spec]
package main

import "fmt"

func main() {
	// In Go there are two distinct categories of types:
	//------------------------
	
	// 1. Value type variables point directly to their value contained in memory.
	// All types explored in previous sections except "pointer", "slice", "map",
	// and "channel" are value types.

	// 2. Reference types variable "pointer, slice, map, and channel", contains the
	// address of the memory location where the value is stored.

	// Merory allocation in Go
	//------------------------
	// Go has two allocation primitives, the built-in functions new and make. 
	// They do different things and apply to different types, which can be 
	// confusing, but the rules are simple.

	// allocation with new
	//------------------------
	// Let's talk about new first. It's a built-in function that allocates memory,
	// but unlike its namesakes in some other languages it does not initialize
	// the memory, it only zeros it. That is, new(T) allocates zeroed storage for
	// a new item of type T and returns its address, a value of type *T.
	// In Go terminology, it returns a pointer to a newly allocated zero value of type T.

	// Since the memory returned by new is zeroed, it's helpful to arrange when
	// designing your data structures that the zero value of each type can be used without
	// further initialization.
	
	// Memory allocated as a result of declaring a variable of value type is zeroed. This is
	// know as the zero value of the type. This behavior is illustated in the following
	// code fragment

	var (
		a uint8
		b int
		c float64
		d bool
		e string
		f complex128
		g [4]bool
	)

	type car struct {
		make  string
		model string
		year  int16
	}

	h := car{}

	fmt.Printf("a = %d (%T)\n", a, a)
	fmt.Printf("b = %d (%T)\n", b, b)
	fmt.Printf("c = %f (%T)\n", c, c)
	fmt.Printf("d = %t (%T)\n", d, d)
	fmt.Printf("e = %#v (%T)\n", e, e)
	fmt.Printf("f = %f (%T)\n", f, f)
	fmt.Printf("g = %#v (%T)\n", g, g)
	fmt.Printf("h = %#v (%T)\n", h, h)

	// allocation with make
	//------------------------
	// The built-in function make(T, args) serves a purpose different from new(T).
	// It creates slices, maps, and channels only, and it returns an initialized (not zeroed)
	// value of type T (not *T). The reason for the distinction is that these three types
	// represent, under the covers, references to data structures that must be initialized 
	// before use. A slice, for example, is a three-item descriptor containing a pointer to 
	// the data (inside an array), the length, and the capacity, and until those items are
	// initialized, the slice is nil. For slices, maps, and channels, make initializes the
	// internal data structure and prepares the value for use.

	var i []int64
	if i == nil {
		fmt.Println("i is nil")
	}
	fmt.Printf("i = %#v (%T)\n", i, i)

	j := make([]int64, 5)
	if j == nil {
		fmt.Println("j is nil")
	}
	fmt.Printf("j = %#v (%T)\n", j, j)

	var k map[string]int64
	if k == nil {
		fmt.Println("k is nil")
	}
	fmt.Printf("k = %#v (%T)\n", k, k)

	l := make(map[string]int64, 5)
	fmt.Printf("l = %#v (%T)\n", l, l)

	var m chan string
	if m == nil {
		fmt.Println("m is nil")
	}
	fmt.Printf("m = %#v (%T)\n", m, m)

	n := make(chan string)
	fmt.Printf("n = %#v (%T)\n", n, n)
}
