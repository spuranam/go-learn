package main

import "fmt"

func main() {
	// The type *T is a pointer to a T value. Its zero value is nil.
	// pointer's value refers directly to (or "points to") another value
	// stored elsewhere in the computer memory using its address.
	// NOTE: Value of an uninitialized pointer is nil.
	// NOTE: Go does not support pointer arithmatic like C or C++

	// In Go a pointer is represented using the * (asterisk) character
	// followed by the type of the stored value
	var p *int             //declares a pointer to an int
	fmt.Printf("%#v\n", p) // p is unintialized (no memory has been allocated) hence its value is nil.

	// Trying derefernce a nil pointer will lead to run-time error and
	// your program will crash.
	// to see this in action uncomment the next line
	//*p = 24

	// its easy to check if pointer is nil
	if p == nil {
		fmt.Printf("p is a nil pointer\n")
	}

	p = new(int) // p is initailized to zero value of int, 0 (zero)
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

	// In practice its very rare that you want to use pointers to primitive
	// types, or for that matter referrence types.
	// most often they are used to point to the values of type struct.
	// For more detailed discussion see this excellent blog post:
	// http://www.goinggo.net/2014/12/using-pointers-in-go.html

}
