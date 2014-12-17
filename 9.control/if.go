package main

import (
	"fmt"
	"math"
)

func main() {

	num := 16.0
	upperBound := 10.0

	// The syntax is also slightly different: there are no parentheses and
	// the bodies must always be brace-delimited.
	if num > upperBound {
		fmt.Printf("number is %f\n", num)
	}

	// If with a short statement.
	// NOTE: the scope of variable sqr is limited the enclosing if statement
	// body
	if sqr := math.Sqrt(num); sqr < upperBound {
		fmt.Printf("Square root of %f is %f\n", num, sqr)
	}

	// uncomment this line to verify that teh variable sqr is not visible
	// outside the previous defined if block
	//fmt.Printf("square root of .2f is %.2f\n", sqr, num)

	// NOTE that its idiomatic to write an if-else-if-else chain as a
	// switch instead.
	if num < upperBound {
		fmt.Printf("%.2f < %.2f\n", num, upperBound)
	} else if num == upperBound {
		fmt.Printf("%.2f == %.2f\n", num, upperBound)
	} else {
		fmt.Printf("%.2f > %.2f\n", num, upperBound)
	}
}
