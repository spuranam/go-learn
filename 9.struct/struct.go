package main

import "fmt"

// A struct is a sequence of named elements, called fields, each of which has a
// name and a type. Field names may be specified explicitly (IdentifierList) or
// implicitly (AnonymousField). Within a struct, non-blank field names must be unique.

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

// By not specifying a named field for Person, this is called an anonymous field.
// Such anonymous fields allows for embedding.
// Go does not provide the typical, type-driven notion of subclassing, but it
// does have the ability to “borrow” pieces of an implementation by embedding
// types within a struct or interface. More on this in later lessons
type Citizen struct {
	Country string
	Person
}

// An empty struct, is just like any other struct, except it does not have any
// fields defined. So what can it be used for. A channel of the type struct{} is
// used for signaling between go routinues, as we will learn in a later lesson
// see the folowing blog post for more relavent discussion on this topic
// http://dave.cheney.net/2014/03/25/the-empty-struct
type S struct{}

func main() {
	// composite literal syntax
	p := Person{"John Doe", Address{"1234", "South Street", "Dearbon", "Michigan", "38123"}}
	fmt.Printf("Person p details:%#v\n", p)

	// its possible to assign values to indivigual fields. Fields are
	// to their respective data types zero value
	var p1 Person
	p1.Name = "Marry Jane"
	p1.Address.Number = "8799"
	p1.Address.Street = "Nightingale Ave"
	p1.Address.City = "Florence"
	// notice the value of the state and zip fields
	fmt.Printf("Person p1 details %#v\n", p1)

	p2 := Person{Name: "Harry Truman"}
	fmt.Printf("Partial struct %#v\n", p2)

	// structs are mutabale
	p2.Name = "Abraham Lincoln"
	fmt.Printf("Mutated struct %#v\n", p2)

	// access the struct fileds
	fmt.Printf("Hello %s how are you?\n", p.Name)

	// The special prefix & returns a pointer to the struct value.
	p3 := &p1
	fmt.Printf("struct pointer %#v\n", p3)

	// The indirection through the pointer is transparent, i.e. we dont have
	// dereference the pointer, this happens automatically
	fmt.Printf("person p3 is %s\n", p3.Name)

}
