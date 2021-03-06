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
	// Trying to access or set an index that does not exists, results in compile-time error,
	// since go compiler does array bound ckecking
	// uncomment the following line to see the error in action
	//numbers[3] = 98

	// array elements are zeroed .i.e. the elements that are not assigned a value,
	// will intialized to their respective types zero value. For example in
	// case of numbers array the value of third (3) element would set to zero (0)
	fmt.Printf("numbers = %#v\n", numbers)

	// declare an array using the literal syntax
	carModels := [3]string{"Fusion", "Fiesta", "Mustang"}
	fmt.Printf("carModels = %#v\n", carModels)

	// you can use an ellipsis to specify an implicit length when you pass the values
	// compiler would count the number elements and fill in the ellipsis
	weekDays := [...]string{"Mon", "Tue", "Wed", "Thus", "Fri", "Sat", "Sun"}
	fmt.Printf("Days in week %#v\n", weekDays)

	// you can determine the length using the len function
	fmt.Printf("Number of elements in carModels %d\n", len(carModels))

	// use the subscript notation to access an array element
	fmt.Printf("First day of the week is %s\n", weekDays[0])

	// you can also iterate over the array using for..range loop
	for indx, val := range carModels {
		fmt.Printf("car at index %d is %s\n", indx, val)
	}

	// or using for loop
	for m := 0; m < len(carModels); m++ {
		fmt.Printf("For: element at index %d is %s\n", m, carModels[m])
	}

	// multi-dimensional arrays
	var arr [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arr[i][j] = fmt.Sprintf("r%d-c%d", i+1, j+1)
		}
	}
	fmt.Printf("%q\n", arr)

	// multi-dimensional arrays, literal syntax
	farr := [2][4]float64{{1.1, 1.2, 1.3, 1.4}, {2.1, 2.2, 2.3, 2.4}}
	fmt.Printf("%#v\n", farr)

	// size of multi-dimensional arrays
	var a [3][5]int8
	for i, h := range a {
		fmt.Printf("Number of columns:%d in array %d\n", len(h), i) // each one prints 5
	}
	fmt.Printf("Number of rows:%d\n", len(a)) // prints 3, the length of the external array

}
