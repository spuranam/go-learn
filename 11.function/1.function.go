package main

import (
	"errors"
	"fmt"
)

// Functions main purpose is to break a large program into a number of smaller
// tasks (functions). It also helps enforce the D-R-Y (dont repeat yourself)
// principle, the same task can be invoked several times, so a function promotes code reuse.

// There are 3 types of functions in Go:
// - Normal functions with an identifier
// - Anonymous or lamda functions
// - Methods

// Any of these can have parameters and return values. The definition of all the
// function parameters and return values, together with their types, is called
// the function signature.

// The function main is special, go programs begins execution in the function
// named main located in package main.

func main() {

	// You invoke the function by specifying the name of the package it is
	// defined in followed by a . (dot) followed by the name of the function
	// and any parameters within a set of parantheses.
	// However if the function is defined in the current package then it can
	// invoked by referring to its name and providing the any parameters
	Greet()

	x, y := 1, 2
	// You can capture the returned value either assigning it to a variable or
	// passing it as argument to another function
	r := add(x, y)
	fmt.Printf("The sum of %d and %d is %d\n", x, y, r)

	numr, denom := 3, 0
	quot, err := divide(numr, denom)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result of %d / %d is %d\n", numr, denom, quot)
	}

	// One or more returned values can be discarded using the blank identifier _ (underscore)

	// When would this be useful?
	// Recall that its illegal to declare a variable and not use it. Now suppose
	// you have function that you wrote or calling a function defined in some
	// package, but you do not want to consume one or more returned values.
	// The only way out, is to use the blank indentifier.

	// For example here the second retured value error is being discarded
	// NOTE: Its considered bad parctice to ignore errors. Please dont do this in normal code.
	// The pupose is to demonstrate how to discard one or more returned values.
	q, _ := divide(numr, denom)
	fmt.Printf("Result of %d / %d is %d\n", numr, denom, q)

	// call the variadic function sum
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// n... is the shorthand to pass each element of the slice to the function
	s := sum("Sum of input number is:", n...)
	fmt.Printf("%s", s)

	// call recursive function
	fmt.Printf("Fibonacci of 10 is %d\n", fibonacci(10))

	// A function that does not have a name is called anonymous function
	// (also known under the names of a lambda function, a function literal, or a closure).
	// Such a function cannot stand on its own, since complier would throw an error.
	// Hence such functions  must be either assigned to variable or must be
	// directly invoked or returned as output value of a named function

	// The variable prod's value is a function that returns product of the passed numbers
	prod := func(x, y int) int {
		return x * y
	}
	fmt.Printf("%d * %d = %d\n", x, y, prod(x, y))

	// To directly invoke an anonymous function, put pair of () after the closing
	// '}' brackets within which we can list the optional parameters to the function.
	func(x, y int) {
		fmt.Printf("%d %% %d = %d\n", x, y, x%y)
	}(4, 2) // 4,2 are the parameters to the anonymous function
}

// The keyword func introduces a function
// A function can take zero or more paramaters
// One of the functions parameters can accept arbitary number of values
// A function can return zero or more values, unlike other languages
// A function can take other function as paramters or return functions as return values
// A function which starts with capital/uppercase letter is exported (visible/accessible from other packages)
// The order in which functions are defined in go source file is of no
// consequence, however it is idiomatic to define the main function as first function.

// Greet is a function that take no paraters and returns no values. Notice this
// function start with capital letter hence it is said to be a exported function.
func Greet() {
	fmt.Printf("Hello how are you?\n")
}

// add is a function that take two parameters of type int and return a single
// value of the type int. Notice this function start with lowercase letter hence
// it is said to be un-exported function.
func add(x, y int) int {
	return x + y
}

var ErrDivideByZero = errors.New("Division by zero is not allowed")

// A function can return one or more values. This forms very foundations of Go's error
// handling machinary, since unlike other language go does not support exception handling
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivideByZero
	}
	return x / y, nil
}

// A go function can accept arbitary number of values for a paramater.
// Such functions are called as variadic functions. A variadic parameter
// has ... prefixed to its type.
// There can only be one such parameter. If a function happens to take
// more than one parameter, then the parameter that accepts multiple values
// must be the last one. The following example show one such function.
// The value of the variadic parameter is accessible as a slice within the
// function body, which are be iterated over using the for..range loop
func sum(title string, nums ...int) string {
	fmt.Printf("Data type of nums is %T\n", nums)
	s := 0
	for _, v := range nums {
		s += v
	}
	return fmt.Sprintf("%s %d\n", title, s)
}

// A function that calls itself in the body of the function is called a recursive function.
func fibonacci(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	}

	return fibonacci(num-1) + fibonacci(num-2)
}
