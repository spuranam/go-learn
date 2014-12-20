package main

import "fmt"

func main() {

	// map[T]M is map with element(value) type M and key type T.
	// A map is an unordered group of elements of one type, called the
	// element or value type, indexed by a set of unique keys of another type,
	// called the key type. The value of an uninitialized map is nil.
	// The key can be of any type for which the equality operator is defined,
	// such as integers, floating point and complex numbers, strings, pointers,
	// interfaces (as long as the dynamic type supports equality), structs
	// and arrays. Slices cannot be used as map keys, because equality is
	// not defined on them. Like slices, maps hold references to an underlying
	// data structure.

	// some languages refer to map as hashmap, dictionary

	// some note worthy properties of map are;
	// map is referrence type, like slice
	// zero value of map is nil
	// map's keys are unique
	// deleting non-existing keys do not raise errors
	// creating a duplicate key silently overwrites the existing key
	// map is not safe for use by concurrent functions, extra precautions
	// such as locking must be taken.

	// like slice, maps are initialzed using the make builtin function or by
	// using the composite literal syntax, with colon separated key-value pairs
	// the make function takes type and an optional capacity as parameters

	priceList := make(map[string]float64, 5)
	priceList["eggs"] = 10.20
	priceList["milk"] = 2.45
	fmt.Printf("%#v\n", priceList)

	// distance from sun in astronomical units
	distanceFromSun := map[string]float64{
		"Mercury": 0.387,
		"Venus":   0.722,
		"Earth":   1.00,
		"Mars":    1.52,
	}
	fmt.Printf("Distance from sun in (AU) %#v\n", distanceFromSun)

	// delete built in function can be used delete the items froms the map
	fmt.Printf("Before Delete, %#v\n", priceList)
	delete(priceList, "eggs")
	fmt.Printf("After Delete, %#v\n", priceList)

	// deleting the same item has no effect
	delete(priceList, "eggs")
	fmt.Printf("Delete again %#v\n", priceList)

	// test for presence of key in the map
	milk := priceList["milk"]
	fmt.Printf("price of milk is %f\n", milk)

	// if the key does not exists then you will get back zero value for the value type
	egg := priceList["eggs"]
	fmt.Printf("price of eggs is %f\n", egg)

	// to distinguish between a missing entry and zero value, we use the comma ok idiom
	egg, ok := priceList["eggs"]
	if ok {
		fmt.Printf("price of egg is %f\n", egg)
	} else {
		fmt.Printf("no eggs found in the price list\n")
	}

	// if we add duplicate key then, go would silently overwrite the previous value
	fmt.Printf("Before add %#v\n", priceList)
	priceList["milk"] = 100.00
	fmt.Printf("After add %#v\n", priceList)

	// length of the map can be discovered using the len function
	fmt.Printf("priceList length=%d\n", len(priceList))

	// you can iterate over the map using for..range loop
	for k, v := range distanceFromSun {
		fmt.Printf("key=%s value=%f\n", k, v)
	}
}
