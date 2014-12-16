package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// in there are following diffrent data types
	// 1. boolean types
	var isLeapYear bool = true
	fmt.Printf("%t (%T)\n", isLeapYear, isLeapYear)

	// 2. numeric types
	// unsigned integers "unit" of size 8,16,32,64 bits
	var age uint8 = 5
	fmt.Printf("%d (%T)\n", age, age)

	// signed integers "int" of size 8,16,32,64 bits
	var num int8 = -50
	fmt.Printf("%d, (%T)\n", num, num)

	// floating point number "float" of size 32, 64 bits
	var detriotToChicago float64 = 308.60
	fmt.Printf("%.2f (%T)\n", detriotToChicago, detriotToChicago)

	// byte is an alias for uint8
	var fileSize byte = 8
	fmt.Printf("%d (%T)\n", fileSize, fileSize)

	// rune is an alias for int32, used to represent unicode code points
	var letter rune = 104 //ascii character 'h'
	fmt.Printf("%#U (%T)\n", letter, letter)

	// complex number "complex" with 32,64 bits real and imaginary parts
	var z complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("%.4f (%T)\n", z, z)

	// string
	var name string = "John Doe"
	fmt.Printf("%s (%T)\n", name, name)

	// array are numbered sequence of elements of single type.
	// note the size of array is part of the type
	var names [2]string
	names[0] = "John Doe"
	names[1] = "Marry Jane"
	fmt.Printf("%v (%T)\n", names, names)

	// array literals syntax
	// the following two definitions are identical
	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday", "Sunday"}

	// the ... causes the compiler to count the number of elements in the array
	//days := [...]string{"Monday","Tuesday","Wednesday","Thrusday","Friday","Saturday","Sunday"}
	fmt.Printf("%v (%T)\n", days, days)

	// slice is descriptor of the underlying array. Like array's slice is also a numbered
	// sequence of elements of single type, but unlike array whoes size is fixed and cannot be
	// altered, silce can grow and shrink. A slice is a descriptor of an array segment.
	// It consists of a pointer to the array, the length of the segment, and
	// its capacity (the maximum length of the segment).

	// slice literal synatx
	carParts := []string{"engine", "wheels", "chasis", "transmission"}
	fmt.Printf("%#v (%T)\n", carParts, carParts)

	// make([]T,len,cap) if cap is omitted the it would be equal to length.
	// the make built in function creates the array, and returns the slice that point to it.
	grains := make([]string, 5, 5)
	grains[0] = "rice"
	grains[1] = "wheat"
	fmt.Printf("%#v (%T)\n", grains, grains)

	// struct is composite type that have sequence of named elements called fields.
	// each field has a name and type. field name can be omitted this is called embeding
	type car struct {
		make  string
		model string
		year  int16
	}

	c := car{"Ford", "Fusion", 2015}
	fmt.Printf("%#v (%T)\n", c, c)

	c1 := car{model: "Escape", make: "Ford"}
	fmt.Printf("%#v (%T)\n", c1, c1)

	//pointer

	//function

	//interface

	//map

	//channel
}
