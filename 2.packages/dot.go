package main

import . "./calc"
import "fmt"

func main() {
	x, y := 6, 3

	fmt.Printf("Sum of %d and %d is %d\n", x, y, Add(x, y))
	fmt.Printf("Diffrence of %d and %d is %d\n", x, y, Substract(x, y))
}
