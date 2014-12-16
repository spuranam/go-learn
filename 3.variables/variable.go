package main

import "fmt"

// declare a variable, notice the use of "var" keyword
// un-initialized variable "City" is initialized to the zero value, in this case an empty string
// indentifiers that start with uppercase letter are exported .i.e. they are visible outside the
// package they are defined.
// addtionally the type of variable appears after the name, unlike Java, C, or C++
var City string

var a int
var c float64
var d bool

// you can group several variable, saves you from the repeating the "var" keyword
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

	// this short hand for declaring and initializng the variable is only available within
	// function bodies
	name := "Penguin"
	canFly, avgLifespan := false, 15

	fmt.Println(name, canFly, avgLifespan)
}
