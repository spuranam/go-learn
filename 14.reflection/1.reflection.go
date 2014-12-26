package main

import (
	"fmt"
	"reflect"
)

// Reflection is the ability of a program to examine its own structure, particularly
// through the types; it’s a form of metaprogramming. Reflection can be used to
// investigate types and variables at runtime, e.g. its size, its methods, and it
// can also call these methods ‘dynamically’.

// It’s a powerful tool that should be used with care and avoided unless strictly necessary.
// The standard library package "refelct" houses Go's reflection capabilities.

// Further reading:
// http://blog.golang.org/laws-of-reflection

var area float64 = 32.98

type MyInt int

var x MyInt = 14

var y uint8 = 'x'

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("type:", reflect.TypeOf(area))
	v := reflect.ValueOf(area)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("The kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("float:", v.Float()) //grab the float value
	fmt.Println("iface:", v.Interface())
	fmt.Println("Settability of v:", v.CanSet())
	fmt.Println("=====================")

	// methods of Value operate on the largest type that can hold the value:
	// int64 for all the signed integers, for instance
	v = reflect.ValueOf(y)
	fmt.Println("type:", v.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	y = uint8(v.Uint())                                       // v.Uint returns a uint64.
	fmt.Println("=====================")

	// Kind of a reflection object describes the underlying type, not the static type
	fmt.Println("type:", reflect.TypeOf(x))
	v = reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("The kind is int:", v.Kind() == reflect.Int)
	fmt.Println("float:", v.Int()) //grab the int value
	fmt.Println("iface:", v.Interface())

	// Settability is a bit like addressability, but stricter. It's the
	// property that a reflection object can modify the actual storage that
	// was used to create the reflection object. Settability is determined
	// by whether the reflection object holds the original item.

	v = reflect.ValueOf(area)
	fmt.Println("Setability:", v.CanSet())

	//this will panic
	//v.SetFloat(1.23)

	// The reason is that a copy of the variable area is passes to the
	// function reflect.ValueOf , so the interface value created as the
	// argument to reflect.ValueOf is a copy of area, not area itself. Thus,
	// if the statement v.SetFloat(1.23) were allowed to succeed, it would
	// not update area, even though v looks like it was created from area.
	// Instead, it would update the copy of area stored inside the reflection
	// value and x itself would be unaffected. That would be confusing and
	// useless, so it is illegal, and settability is the property used to
	// avoid this issue.

	// If were to pass area to function like f(area), we can't expect f to
	// modify the value of area, because we passed a copy of area's value
	// not the area itself. If we want to f to modify the area then we pass
	// the address of the area to f, using the address of operator f(&area)

	// Now if want to modify the value of area using the reflection, we must
	// refelction libarary a pointer to the value we want to modify.

	p := reflect.ValueOf(&area)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settablity of p:", p.CanSet())
	// p is pointer, we have to dereference p before we can set its value
	// to get what p points to we can use Elem method of value
	v = p.Elem()
	fmt.Println("Settability of v:", v.CanSet())
	v.SetFloat(1.234)
	fmt.Println("v:", v.Interface())
	fmt.Printf("area: %2f\n", area)

	// structs and reflection
	p1 := Person{"John Doe", 23}
	pp := reflect.ValueOf(&p1).Elem()

	for i := 0; i < pp.NumField(); i++ {
		fmt.Printf("%d: %s %s = %v\n",
			i, //field index
			pp.Type().Field(i).Name, //field name
			pp.Field(i).Type(),      //field type
			pp.Field(i).Interface()) //field value
	}

	// set the struct field value
	pp.Field(0).SetString("Marry Jane")
	pp.Field(1).SetInt(90)
	fmt.Printf("%#v\n", p1)
}
