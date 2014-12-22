// package is a collection types, variables, constants, functions defined in
// one or more files. Some of these identifiers are exported, so that other
// packages can import them.

// Programs starts executing in the main package, specifically in the main function
// NOTE: Its complie time error to import a package and not use it. The only exception
// to this to use _ (underscore) import statement. see sample se.go for details.
package main

// Import declaration declares library packages referenced in this file.
import "./calc" // a user defined package
import "fmt" // A package in the Go standard library.

func main() {
	x, y := 6, 3

	fmt.Printf("Sum of %d and %d is %d\n", x, y, calc.Add(x, y))
	fmt.Printf("Diffrence of %d and %d is %d\n", x, y, calc.Substract(x, y))
}
