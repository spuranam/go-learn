package main

import "fmt"

type automobile struct {
	make  string
	model string
	year  int
	miles float64
	color string
}

// In Go methods reciever can be either a value of type or pointer to value.
// If we call a method that takes a value, on a pointer, Go is smart enough to
// dereference the pointer and pass the underlying value as the method’s receiver.

// NOTE: The rule about pointers vs. values for receivers is that value methods can be
// invoked on pointers and values, but pointer methods can only be invoked on pointers.
// This rule arises because pointer methods can modify the receiver; invoking
// them on a value would cause the method to receive a copy of the value, so any
// modifications would be discarded. The language therefore disallows this mistake.
// There is a handy exception, though. When the value is addressable (i.e., it is a
// variable, a dereferenced pointer, an array or slice item, or an addressable field
// in a struct), the language  takes care of the common case of invoking a pointer
// method on a value by inserting the address operator automatically.

// NOTE: interface methods are not addressable, hence the above exception will not
// apply if the method(s) of a type satisfy an interface. see the topic intrefaces
// for more relavent details.

// The methods start, stop and honk work either way. It only reads a. It does'nt
// wether it is reading the original value (through a pointer) or a copy of the value
func (a automobile) start() {
	fmt.Printf("%s %s has started and is ready to go\n", a.make, a.model)
}

func (a automobile) stop() {
	fmt.Printf("%s %s has come to stop\n", a.make, a.model)

}

func (a automobile) honk() {
	fmt.Printf("%s %s is honking its horn\n", a.make, a.model)
}

// The paint method has no effect when a is automobile. paint mutates a. When a is
// a value (non-pointer) type, the method sees a copy of automobile and cannot mutate
// the original value
func (a automobile) paint(c string) {
	a.color = c
	fmt.Printf("%s %s has been painted %s\n", a.make, a.model, c)
}

// There are two reasons to use a pointer receiver. First, to avoid copying the
// value on each method call (more efficient if the value type is a large struct).
// Second, so that the method can modify the value that its receiver points to.
func (a *automobile) resetMiles() {
	a.miles = 0.0
}

func main() {
	fusion := automobile{"Ford", "Fusion", 2014, 44.0, "White"}
	mustang := automobile{"Ford", "Mustang", 2010, 12378.45, "Red"}

	fusion.start()
	fusion.honk()
	fusion.stop()

	mustang.start()
	mustang.honk()
	mustang.stop()

	fmt.Printf("%s %s is of %s color\n", mustang.make, mustang.model, mustang.color)
	// mutate mustang
	mustang.paint("Black")
	// since paint methods reciever type is defined as value of automobile
	// the paint method is pass a copy of automobile, which it mutates. But the
	// original value is left unchnaged
	fmt.Printf("The color of %s %s is %s\n", mustang.make, mustang.model, mustang.color)

	fmt.Printf("%s %s has %.2f miles on it\n", mustang.make, mustang.model, mustang.miles)
	// since the value mustang is addressable Go would automatically insert (prefeix)
	// the address operator to the value mustang, hence the following two lines
	// are equivalent
	mustang.resetMiles()
	//(&mustang).resetMiles() // same as previus method call
	fmt.Printf("%s %s has %.2f miles on it\n", mustang.make, mustang.model, mustang.miles)
}
