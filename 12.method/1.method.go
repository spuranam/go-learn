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
