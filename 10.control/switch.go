package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	num := 16.0
	upperBound := 10.0

	// A case body breaks automatically, unless it ends with a fallthrough statement.
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

	// switch with no condition
	switch {
	case num < upperBound:
		fmt.Printf("%.2f < %.2f\n", num, upperBound)
	case num == upperBound:
		fmt.Printf("%.2f == %.2f\n", num, upperBound)
	default:
		fmt.Printf("%.2f > %.2f\n", num, upperBound)
	}

	// switch with condition and a statement
	// NOTE the scope of variable oe is limited to the switch block

	// seed the random number generator
	rand.Seed(int64(time.Now().Nanosecond()))
	// generate a random number between 1 and 100
	rnum := rand.Intn(100)

	switch oe := rnum % 2; oe {
	case 0:
		fmt.Printf("Number %d is even\n", rnum)
	case 1:
		fmt.Printf("Number %d is odd\n", rnum)
	}

	// cases can be composed of compound statements
	switch {
	case rnum%2 == 0 && rnum%4 == 0:
		fmt.Printf("The number %d is divisible by 2 and 4\n", rnum)
	case rnum%2 == 0 || rnum%4 == 0:
		fmt.Printf("The number is %d is not divisible by both 2 and 4\n", rnum)
	default:
		fmt.Printf("I do know what to do with %d\n", rnum)
	}

	// if want switch to continue evalution after the first match you
	// do so with fallthrough statement.
	// NOTE: you can not fallthrough the final case in switch
	switch {
	case rnum%2 == 0 && rnum%4 == 0:
		fmt.Printf("The number %d is divisible by 2 and 4\n", rnum)
		fallthrough
	case rnum%3 == 0:
		fmt.Printf("The number is %d is divisible by both 3\n", rnum)
		fallthrough
	default:
		fmt.Printf("I do know what to do with %d\n", rnum)
	}

	// A break statement can be used to terminate the switch early
	// Sometimes, though, it's necessary to break out of a surrounding loop,
	// not the switch, and in Go that can be accomplished by putting a label
	// on the loop and "breaking" to that label. This example shows both uses.
	// addtionally this example also shows that cases can evaluate multiple
	// values at once
LOOP:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6, 8:
			fmt.Printf("Even number %d\n", i)
		case 3, 5, 7, 9:
			break LOOP
		default:
			fmt.Printf("What shall we do with %d\n", i)
		}
	}

}
