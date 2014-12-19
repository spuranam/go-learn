package main

import "fmt"

func main() {

	// Go has only one loop control statement. for statement is used to run a
	// block of code repeatedly. The iteration is controlled by a condition,
	// a "for" clause, or a "range" clause.

	// In its simplest form, a "for" statement specifies the repeated execution
	// of a block as long as a boolean condition evaluates to true. The
	// condition is evaluated before each iteration. If the condition is absent,
	// it is equivalent to the boolean value true.
	// NOTE: this is equivalent to a while loop in other languages.
	num := 0
	for num < 3 {
		fmt.Printf("simple for loop id %d\n", num)
		num++
	}

	// A "for" statement with a ForClause is also controlled by its condition,
	// but additionally it may specify an init and a post statement, such as
	// an assignment, an increment or decrement statement. The init statement
	// may be a short variable declaration, but the post statement must not.
	// Variables declared by the init statement are re-used in each iteration
	for i := 0; i < 4; i++ {
		fmt.Printf("loop with ForClause %d\n", i)
	}

	// if the condition statement along with init and post statement are ommited
	// then for statement will bahave like a infinite loop.
	// for statement can be terminated early using the break statement.
	// the current iteration can skipped using the continue statement.
	// The example shows all three uses.
	i, s := 0, 0

	for {
		i += 1
		if i >= 10 {
			break
		}
		if i%2 == 0 {
			continue
		}
		s += i
	}
	fmt.Printf("Sum of odd numbers less than 10 is %d\n", s)

	// A "for" statement with a "range" clause iterates through all entries
	// of an array, slice, string or map, or values received on a channel.
	// For each entry it assigns iteration values to corresponding iteration
	// variables if present and then executes the block.
	a := [4]int{2, 34, 56, 78}

	for i, v := range a {
		fmt.Printf("Index=%d Element=%d\n", i, v)
	}

	// If ommit the index, then the blank identifier _ (underscore) can be used
	// to discard the index
	for _, v := range a {
		fmt.Printf("Array element %d\n", v)
	}

	m := map[string]int{"Fusion": 30200, "Mustang": 45678}
	for k, v := range m {
		fmt.Printf("Key=%s value=%d\n", k, v)
	}
}
