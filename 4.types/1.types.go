// Attributions
// some of the details below have been reproduced here from;
// Effective Go [http://golang.org/doc/effective_go.html]
// The Go Programming Language Specification [http://golang.org/ref/spec]
package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"strings"
	"time"
)

// interface defines a set of methods. A type is said to satify an interface
// when it implements all methods listed in the interface type.

// all types in go satisfy and empty interface "interface{}", since an empty
// interface contains no methods. A alluded to ealier a type satfies an interface
// if it implements all the methods called for, since an empty interface does
// define and methods hence all type satisfy an empty interface.

type Upcaser interface {
	upcase() string
}

type greetings string

func (g greetings) upcase() string {
	return strings.ToUpper(string(g))
}

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

	// pointers value refers directly to (or "points to") another value
	// stored elsewhere in the computer memory using its address
	// To get the address of a value use address-of operator "&".
	// NOTE: Value of an uninitialized pointer is nil.
	// NOTE: Go does not support pointer arithmatic like C or C++
	cp := &c1
	fmt.Printf("%#v (%T)\n", cp, cp)

	// function type denotes the set of all function with same paramter and result type
	// functions are declared using the "func" keyword
	// attribution:
	// the following example is slightly modified reproduction of:
	// http://jordanorelli.com/post/42369331748/function-types-in-go-golang
	type binFunc func(x, y int) int

	// seed the random number generator
	rand.Seed(int64(time.Now().Nanosecond()))

	// create a slice of function of the type binFunc
	fns := []binFunc{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x * y },
		func(x, y int) int { return x / y },
		func(x, y int) int { return x % y },
	}

	// randomly pick one function from the slice
	f := fns[rand.Intn(len(fns))]
	fmt.Printf("%v (%T)\n", f, f)

	n, m := 1, 2
	fmt.Printf("%v (%T)\n", f(n, m), f(n, m))

	greet := greetings("hello")
	gu := greet.upcase()
	fmt.Printf("%s (%T)\n", gu, gu)

	// A map is an unordered group of elements of one type, called the
	// element or value type, indexed by a set of unique keys of another type,
	// called the key type. The value of an uninitialized map is nil.
	// The key can be of any type for which the equality operator is defined,
	// such as integers, floating point and complex numbers, strings, pointers,
	// interfaces (as long as the dynamic type supports equality), structs
	// and arrays. Slices cannot be used as map keys, because equality is
	// not defined on them. Like slices, maps hold references to an underlying
	// data structure.

	unitPrice := make(map[string]float64)
	unitPrice["egg"] = 10.20
	unitPrice["milk"] = 2.45
	fmt.Printf("%#v (%T)\n", unitPrice, unitPrice)

	polygons := map[string]int{"square": 4, "triagle": 3, "pentagon": 5}
	fmt.Printf("%#v (%T)\n", polygons, polygons)

	//channel type provides a means of communications for concurrently executing function.
	// the details of which will be covered in its own section later in the course.
}
