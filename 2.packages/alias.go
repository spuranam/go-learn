package main

import c "./calc"
import "fmt"

func main() {
	x, y := 6, 3

	fmt.Printf("Sum of %d and %d is %d\n", x, y, c.Add(x, y))
	fmt.Printf("Diffrence of %d and %d is %d\n", x, y, c.Substract(x, y))
}
