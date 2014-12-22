package main

import "fmt"

// A variable is a storage location, with a specific type and an associated name.
// Names must start with a letter and may contain letters, numbers or the _ (underscore) symbol.
// "var" keyword declare a variable.

// Addtionally the type of variable appears after the name, unlike Java, C, or C++

// variable scope
// variables defined outside of the function are accessible to other functions. For example
// the variable name defined inside the main function will not available in other
// functions. Whereas the variable Car will be.

// Varirable whoes name start with uppercase (capital) letter, are exported 
// .i.e. they are visible outside the package they are defined in.

// NOTE: Its complie time error to define a variable and not use it
var City string

var a int
var c float64
var d bool

// Go has a shorthand when you have to define many variables
var (
	e    int64
	f, g uint32
	b    complex128
)

// declare and initialize several variable of different types
// note the Go complier was able to infer the types from their respective values
var age, height, hairColor, hasCar = 20, 66.0, "black", true

func main() {
	fmt.Printf("City=%s age=%d height=%.2f hairColor=%s hasCar=%t\n", City, age, height, hairColor, hasCar)
	fmt.Printf("City=%T age=%T height=%T hairColor=%T hasCar=%T\n", City, age, height, hairColor, hasCar)
	fmt.Printf("a=%#v b=%#v c=%#v d=%#v e=%#v f=%#v g=%#v\n", a, b, c, d, e, f, g)

	// this shorthand for declaring and initializng the variable is only available within
	// function bodies
	name := "Penguin"
	canFly, avgLifespan := false, 15

	fmt.Println(name, canFly, avgLifespan)
}
