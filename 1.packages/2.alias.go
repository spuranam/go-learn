package main

// import c "./calc" - Creates an alias of calc. Preceed all calc package
// content with c instead of calc.

import c "./calc"
import "fmt"

func main() {
	x, y := 6, 3

	fmt.Printf("Sum of %d and %d is %d\n", x, y, c.Add(x, y))
	fmt.Printf("Diffrence of %d and %d is %d\n", x, y, c.Substract(x, y))
}
