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

	// If with a short statement
	if sqr := math.Sqrt(num); sqr < upperBound {
		fmt.Printf("Square root of %f is %f\n", num, sqr)
	}

	// if and else
	if num < upperBound {
		fmt.Printf("%.2f < %.2f\n", num, upperBound)
	} else if num == upperBound {
		fmt.Printf("%.2f == %.2f\n", num, upperBound)
	} else {
		fmt.Printf("%.2f > %.2f\n", num, upperBound)
	}
}
