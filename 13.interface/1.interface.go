package main

import (
	"fmt"
	"math"
)

// An interface defines behaviour of an object. In go an interface is a type
// that specifies one or more method signatures collectively called method set.
// Interfaces are wholly abstract, so it is not possible to instantiate an interface.
// A variable of interface type can store a value of any type with a method set
// that is any superset of the interface. Such a type is said to implement the
// interface. The value of an uninitialized variable of interface type is nil.

// The interface{} type is the interface that specifies the empty set of methods.
// Every value satisfies the interface{} type whether the value has methods or
// not, if a value does have methods, its set of methods includes the empty set
// of methods as well as the methods it actually has. This is why the interface{}
// type can be used for anyvalue.

// A type doesn’t have to state explicitly that it implements an interface:
// interfaces are satisfied implicitly.
// Multiple types can implement the same interface.
// A type that implements an interface can also have other functions. A type can
// implement many interfaces.
// An interface type can contain a reference to an instance of any of the types
// that implement the interface (an interface has what is called a dynamic type)

type shaper interface {
	area() float64
	perimeter() float64
}

type enlarger interface {
	enlarge(float64)
}

type circle struct {
	radius float64
}

type rectangle struct {
	length, width float64
}

type square struct {
	rectangle //anonymous field (embedding)
}

type triangle struct {
	a, b, c float64 //side of a triangle
}

// The methods area and perimeter on types circle, and rectangle are defined
// with a value reciever since these methods do not mutate (modify) the state of
// their respective types.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// enlarge method is defined with a pointer reciever since it mutates (modifies)
// the passed values state. If the reciver was defined as value then the chnage
// to the state of the value will not visible.
func (c *circle) enlarge(by float64) {
	c.radius += by
}

// notice that we have not defined the area and perimeter methods for the square type
// If an embedded field has methods we can call them on the containing struct,
// and only the embedded field will be passed as the methods’ receiver.
func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r *rectangle) enlarge(by float64) {
	r.length += by
	r.width += by
}

// geometry takes any values that satisfy the shaper interface and is able
// dynamically dispatch the call to correct area and perimeter methods
func geometry(s shaper) {
	fmt.Printf("Area of %#v is %.2f\n", s, s.area())
	fmt.Printf("Perimeter of %#v is %.2f\n", s, s.perimeter())
}

func main() {
	c := circle{2.12}
	r := rectangle{4.3, 3.5}
	s := square{rectangle{4.6, 4.6}}

	geometry(c)
	geometry(r)
	geometry(s)

	fmt.Printf("Before enlarge %#v\n", c)
	// Since the circle type is addressable Go would automatically modify the
	// call to enlarge method by passing a pointer to c i.e. it would modify
	// the method call to read (&c).enlarge(1.1)
	c.enlarge(1.1)
	fmt.Printf("After enlarge %#v\n", c)

	// declare t if of type interface{}
	var t interface{}
	t = triangle{1.0, 2.0, 3.0}

	// we can use type assertion to check at runtime if a type satifies an interface
	_, ok := t.(shaper)
	if !ok {
		fmt.Printf("%#v does not statisfy shaper interface\n", t)
	}
}
