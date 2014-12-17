package main

import (
	"fmt"
	"math"
)

func main() {
	var num int64 = 2034
	radius := 2.0
	area := math.Pi * radius * radius

	// You can convert one basic type to another similar type with possible loss of accuracy using,
	// the expression T(v) converts the value v to the type T.

	// Assignment between values of diffrent type requires explicit type coversion, since go
	// does not support automatic/implicit type conversion unlike other language

	n := float64(num)
	fmt.Printf("%v (%T)\n", n, n)

	intArea := int64(area)
	fmt.Printf("%f(%T) %d(%T)\n", area, area, intArea, intArea)
}
