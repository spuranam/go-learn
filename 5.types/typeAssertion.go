package main

import "fmt"

func main() {

	// converting values of dissimilar types will cause complier errors as shown below:
	// ./typeAssertion.go:10: cannot convert name (type string) to type int64

	// uncomment the next three line to see the error in action

	//var name string = "Marry Jane"
	//n := int64(name)
	//fmt.Printf("%v (%T)\n", n, n)

	// A type assertion takes an interface value and extracts from it a value of
	// the specified explicit type.

	// But if the value does not contain the specified explicit type, the program
	// will crash with a run-time error. Uncomment nest three lines to see the
	// error in action

	//var name interface{} = "Marry Jane"
	//aNumber := name.(int64)
	//fmt.Printf("%v (%T)\n", aNumber, aNumber)

	// To guard against the run-time program crash, use the "comma, ok" idiom to test,
	// safely, whether the value is of the specified explicit type
	var name interface{} = "Marry Jane"
	aNumber, ok := name.(int64)
	if ok {
		fmt.Printf("Integer value is %d\n", aNumber)
	} else {
		fmt.Printf("value is not integer\n")
	}

	// The type assertion doesn’t have to be done on an empty interface.
	// It’s often used when you have a function taking a param of a specific
	// interface but the function inner code behaves differently based on
	// the actual object type. A concrete example in listed in typeSwitch.go

}
