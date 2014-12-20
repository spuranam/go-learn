package main

import "fmt"

func main() {
	// The type *T is a pointer to value T. Its zero value is nil.
	// Pointer's value refers directly to (or "points to") another value
	// stored elsewhere in the computer memory using its address.
	// NOTE: Value of an uninitialized pointer is nil.
	// NOTE: Go does not support pointer arithmatic like C or C++

	// In Go a pointer is represented using the * (asterisk) character
	// followed by the type of the stored value
	var p *int             //declares a pointer to an int
	fmt.Printf("%#v\n", p) // p is unintialized (no memory has been allocated) hence its value is nil.

	// Dereferencing a nil pointer will lead to run-time error and your program will crash.
	// To see this in action uncomment the next line
	//*p = 24

	// It is easy to check if the pointers value is nil
	if p == nil {
		fmt.Printf("p is a nil pointer\n")
	}

	p = new(int) // p is initialized to zero value of int, 0 (zero)
	fmt.Printf("%#v\n", *p)
	*p = 24 // set p
	fmt.Printf("%#v\n", *p)

	n := "Marry Jane"
	// To get pointer to a value use address-of operator "&"
	np := &n // np points to n
	fmt.Printf("%v\n", np)

	// * (asterisk) is also used to “dereference” pointer variables.
	// Dereferencing a pointer gives us access to the value the pointer points to.
	fmt.Printf("%s\n", *np) //read n through pointer np

	*np = "John Doe" //set n through pointer np
	fmt.Printf("%s\n", n)

	// In practice it is very rare that you want to use a pointer to primitive types.
	// The values of reference type consists of header which contains
	// a referrence to data structure and other meta data. These types are
	// designed to be shared, when you pass a reference type to function
	// only the a copy of the header is passed to the function not the underlying
	// data structure, hence there is no benifit use or create a pointer to
	// values of reference type.

	// More often pointers to a value of struct type are created and passed
	// to function to mutate their state.
	// For more detailed discussion see this excellent blog post:
	// http://www.goinggo.net/2014/12/using-pointers-in-go.html

	// Array and Pointers
	var pToArray *[2]int // declare a pointer to an array of two int
	var ArrayOfP [2]*int // declare an array of two pointers to int

	pToArray = new([2]int)
	for i := 0; i < 2; i++ {
		ArrayOfP[i] = new(int)
		*ArrayOfP[i] = i * 2
		(*pToArray)[i] = i + 2
	}
	fmt.Printf("pToArray = %v (%T)\n", pToArray, pToArray)
	fmt.Printf("ArrayOfP = %v (%T)\n", ArrayOfP, ArrayOfP) // we get memory address
	for i := 0; i < len(ArrayOfP); i++ {
		// by dereferencing the each element we can read the value stored
		fmt.Printf("Elements of ArrayOfP = %v (%T)\n", *ArrayOfP[i], *ArrayOfP[i])
	}
}
