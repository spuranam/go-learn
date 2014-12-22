package main

import "fmt"

type BinOp func(x, y int) int

func main() {
	x, y := 2, 3
	f := CreateBinOp("+")(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, f)

	g := intSeq()
	// everytime we call the function value g we get next int value
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", g())
	}
}

// notice the type of return value of the function CreateBinOp
// is the same as the type of the BinOp, hence we can replace this with BinOp.
// Hence line 17 is equivalent to line 16
//func CreateBinOp(op string) func(x, y int) int {
func CreateBinOp(op string) BinOp {
	var r BinOp // notice the type of r

	switch op {
	case "+":
		r = func(x, y int) int { return x + y }
	case "-":
		r = func(x, y int) int { return x - y }
	case "*":
		r = func(x, y int) int { return x * y }
	case "/":
		r = func(x, y int) int { return x / y }
	}

	return r
}

// Function literals are closures: they may refer to variables defined in a
// surrounding function. Those variables are then shared between the surrounding
// function and the function literal, and they survive as long as they are accessible.
// intSeq returns another function, which we define anonymously in the body of intSeq.
// The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int { // anonymous function
		i += 1
		return i
	}
}
