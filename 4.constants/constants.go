package main

import "fmt"

// constants are defined using the keyword "const".
// note that unlike variable, constants can not be defined ussing the
// shorthand ":=".

const PI = 3.14

//like variables, they can grouped togather

const (
	MONTHS     = 12
	daysInWeek = 7
)

// iota, is handy indetifier every time its value is incremented by 1.
// here is an example from time standard package
const (
	January = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(PI, MONTHS, daysInWeek)
	fmt.Printf("Sunday is %d, March is  %d\n", Sunday, March)
}
