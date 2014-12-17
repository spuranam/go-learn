package main

import "fmt"

func main() {
	// array is one of three collection type built into the Go language.
	// The type [n]T is an array of n values of type T.
	// array are numbered (indexed) sequence of elemnts of a single specifed type.
	// some intresting properties of go array's are:
	// an array identifies a contigous memory segment
	// the length of the array is part of the type
	// once defined the size of the array cannot be changed
	// array index are zero based like many other languages

	// declare an array
	var numbers [3]int64
	numbers[0] = 1
	numbers[1] = 23
	// Trying to access or set and index that does not exists, results in compile-time error,
	// since go compiler does array bound ckecking
	// ./array.go:18: invalid array index 2 (out of bounds for 2-element array)
	// uncomment the following line to it in action
	//numbers[2] = 98

	// array elemnts are zeroed .i.e. the elements that are not assigned a value,
	// will intialized to their respective types zero value. For example in
	// case of numbers array the value of third (3) element would set to zero (0)
	fmt.Printf("numbers = %#v\n", numbers)

	// declare an array using the litral syntax
	carModels := [3]string{"Fusion", "Fiesta", "Mustang"}
	fmt.Printf("carModels = %#v\n", carModels)

	// you can use an ellipsis to specify an implicit length when you pass the values
	// compiler would count the number elements and fill in the ellipsis
	weekDays := [...]string{"Mon", "Tue", "Wed", "Thus", "Fri", "Sat", "Sun"}
	fmt.Printf("Days in week %#v\n", weekDays)

	// you can determine the length using the len function
	fmt.Printf("Number of elements in carModels %d\n", len(carModels))

	// multi-dimensional arrays
	var arr [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arr[i][j] = fmt.Sprintf("r%d-c%d", i+1, j+1)
		}
	}
	fmt.Printf("%q\n", arr)
}
