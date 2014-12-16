package main

import "fmt"

func main() {
	// In Go there are two distinct categories of types:

	// 1. Value type variables point directly to their value contained in memory.
	// All types explored in previous sections except "pointer", "slice", "map",
	// and "channel" are value types.

	// 2. Reference types variable "pointer, slice, map, and channel", contains the
	// address of the memory location where the value is stored.

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

	// reference types on the hand needs to be initialized, either using a literal
	// synatx or by using a built in function called "make".
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
