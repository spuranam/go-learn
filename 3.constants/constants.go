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

// Iota is a handy shorthand available in Go that makes defining of enumerated
// constants easy. Iota value is reset to 0 at every constant declaration
// (a statement starting with const) and within a constant declaration it is
// incremented after each line(ConstSpec). If you use iota in different
// expressions in the same line they will all get the same iota value.

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
