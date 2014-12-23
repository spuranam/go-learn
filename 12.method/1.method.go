package main

import "fmt"

// greetings is a user defined type who'es underlying representaion is string.
// NOTE: greetings and string are of different type, you will have to explicitly
// convert them.
type greetings string

// An object is made up of behaviour and data. In Go the value (instance) of the
// type stores the data (properties), and methods describe the types behaviors.
// Go method is a function that acts on variable of a certain type, called the receiver.
// Go reciever corresponds to java and C#'s' "this" keyword.
// A receiver can be any user defined type, in otherwords methods can only be defined on
// user defined types.
// The collection of all the methods on a given type T (or *T) is called the method set of T (or *T)

// A type’s method set is the set of all the methods that can be called on a value of the type.
// If we have a pointer to a value of a custom type, its method set consists of
// all the methods defined for the type—whether they accept a value or a pointer

// If we have a value of a custom type, its method set consists of all those methods
// defined for the type that accept a value receiver—but not those methods that
// accept a pointer receiver. This isn’t as limiting as it sounds, since if we have
// a value we can still call a method that has a pointer receiver and rely on Go
// to pass the value’s address—providing the value is addressable (i.e., it is a
// variable, a dereferenced pointer, an array or slice item, or an addressable field
// in a struct). So, given the call value. Method() where Method() requires a pointer
// and value is an addressable value, Go will treat the code as if we had written (&value). Method().

// greet is a method that can be called on instances of the type greetings
func (g greetings) greet() {
	fmt.Printf("%s\n", string(g))
}

func main() {
	// three distinct instances of greetings
	hindi := greetings("Namaste")
	english := greetings("Hello")
	mandarin := greetings("Nihao")

	// invoke the geet method
	hindi.greet()
	english.greet()
	mandarin.greet()
}
