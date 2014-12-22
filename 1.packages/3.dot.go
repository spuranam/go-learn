package main

// import . "calc" - Allows content of the package to be accessed directly,
// without the need for it to be preceeded with calc.

import . "./calc"
import "fmt"

// NOTE: if the package that you are importing into, for example in this case
// main, already has a function named Add, then go complier will through an error.
// to see this error in action uncomment following three (3) lines
// func Add(x, y int) int {
// 	return x + y
// }

func main() {
	x, y := 6, 3

	fmt.Printf("Sum of %d and %d is %d\n", x, y, Add(x, y))
	fmt.Printf("Diffrence of %d and %d is %d\n", x, y, Substract(x, y))
}
