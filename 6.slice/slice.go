package main

import "fmt"

func main() {
	// []T is a slice with elements of type T.
	// Like array's slice is also a numbered sequence of elements of
	// single type, but unlike array whoes size is fixed and cannot be
	// altered, silce can grow and shrink. A slice is a descriptor of an array segment.
	// It consists of a pointer to the array, the length of the segment, and
	// its capacity (the maximum length of the segment).

	// see for more details
	// Go Slices: usage and internals [http://blog.golang.org/go-slices-usage-and-internals]
	// http://blog.golang.org/slices
	// http://www.goinggo.net/2013/08/understanding-slices-in-go-programming.html

	// some note worthy properties of slices are:
	// slice is a referrence type
	// slice can grow beyond the size defined at the time of declaration
	// slice can sub divided, each of these sub slices refer to the same underlying array
	// zero value of slice is nil
	// a nil slice has length and capacity of zero (0)

	// built in function make is used to intialize a slice, which creates the
	// underlying array and return the a pointer to this array.
	// make([]T,len,cap) if cap is omitted the it would be equal to length.
	s := make([]int32, 3)
	s[0] = 12
	s[1] = 1234

	// Trying to access or set an index that does not exists, results in run-time crash
	// and the program terminates. To get around this limitation use the built in
	// function "append"
	// uncomment the following line to see the error in action
	// s[3] = 98

	// slice elemnts are zeroed .i.e. the elements that are not assigned a value,
	// will intialized to their respective types zero value. For example in
	// case of numbers array the value of third (3) element would set to zero (0)
	fmt.Printf("%#v\n", s)

	// slice literal syntax
	weekDays := []string{"Mon", "Tue", "Wed", "Thus", "Fri", "Sat", "Sun"}
	fmt.Printf("Days in week %#v\n", weekDays)

	// add elements to slice
	s = append(s, 9868)
	fmt.Printf("%#v\n", s)

	// The length is the number of elements referred to by the slice.
	// The capacity is the number of elements in the underlying array
	// (beginning at the element referred to by the slice pointer)
	// The length and capacity of a slice can be inspected using the built-in len and cap functions.
	a := []int32{1, 2}
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	// see how length and capacity changes as we append to the slice
	// adding thrid element would double to slices initial capacity.
	// If the backing array of a is too small to fit all the given values
	// a bigger array will be allocated. The returned slice will point
	// to the newly allocated array.
	a = append(a, 23)
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	a = append(a, 41)
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	a = append(a, 234)
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	a = append(a, 400)
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	// Slices can be re-sliced, creating a new slice value that points to the same array.
	// The expression s[lo:hi], evaluates to a slice of the elements from lo through hi-1, inclusive.
	// create a new slice pointing to same array, but has no elements
	b := a[0:0]
	fmt.Printf("b = %v\n", b)

	// has exactly one element, specifically it has element at the first index
	b = a[0:1]
	fmt.Printf("b = %v\n", b)

	// contains elements 0 - through 2
	b = a[:len(a)-3]
	fmt.Printf("b = %v\n", b)

	// contains elements 2 - 5
	b = a[2:]
	fmt.Printf("b = %v\n", b)

	// contains all elements
	b = a[:]
	fmt.Printf("b = %v\n", b)

	// slice can be copied to create new slice
	// func copy(dst, src []T) int
	// The copy function supports copying between slices of different
	// lengths (it will copy only up to the smaller number of elements).
	// In addition, copy can handle source and destination slices that share
	// the same underlying array, handling overlapping slices correctly.
	newSlice := []int32{111, 222, 333}
	fmt.Printf("before copy newSlice = %v\n", newSlice)
	copy(newSlice, a)
	fmt.Printf("newSlice = %v\n", newSlice)

	// use the subscript notation to access an slice element.
	// NOTE: trying to access a non-existing index will cause run-time crash
	fmt.Printf("First day of the week is %s\n", weekDays[0])

	// you can also iterate over the array using for..range loop
	for indx, val := range a {
		fmt.Printf("Range: element at index %d is %d\n", indx, val)
	}

	// or using for loop
	for i := 0; i < len(a); i++ {
		fmt.Printf("For: element at index %d is %d\n", i, a[i])
	}
}
