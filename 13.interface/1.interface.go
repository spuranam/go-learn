package main

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

type transporter interface {
	speedup(int)
	slowdown(int)
	wheelcount() int
}

type car struct {
	speed  int
	wheels int
	make   string
}

type bike struct {
	speed  int
	wheels int
	make   string
}

type bus struct {
	speed  int
	wheels int
	make   string
}

func main() {
}
