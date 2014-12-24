package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

// Go does not have an exception mechanism, like the try/catch in
// Java or .NET for instance, hence you cannot throw exceptions.

// Go functions and methods can return a single value or multiple values.
// It is conventional in Go to report errors by returning an error value
// (of type error) as the last (or only) value returned by a function or method
// and to always check any returned errors.

// Go has a predefined error interface type: type error interface { Error() string}
// Since error is an interface we can define a new error either by using the function
// errors.new from the errors standard libarary package, or defining a custom type
// that satisfies the error interface.

// see these blog posts for more on this topic:
// http://blog.golang.org/error-handling-and-go
// http://www.goinggo.net/2014/10/error-handling-in-go-part-i.html
// http://www.goinggo.net/2014/11/error-handling-in-go-part-ii.html

// To define a simple error:
var ErrDivideByZero = errors.New("division by zero is not allowed")

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivideByZero
	}
	return x / y, nil
}

// Often we may want to return a more informative string for example the value
// of the wrong parameter inserted in the error message fmt.Errorf() function,
// comes to our rescue
func divide2(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by %d is not allowed", y)
	}
	return x / y, nil
}

// Any type that satisfies an error interface can used as error value. A custom error
// struct type can contain other useful information apart from the error-message.
// For example the os standard library package http://golang.org/src/os/error.go
// defines a custoim type PathError to provide addtional context of an error

//type PathError struct {
//	Op   string
//	Path string
//	Err  error
//}

//func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

// Opening a file for example may trigger several error-conditions, we can test with
// type assertion or type switch for the exact error, and possibly try a remedy
// or a recovery of the error-situation:
func readFile() {
	_, err := os.Open("UNKNOWN_FILE")
	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("%s %s %s\n", e.Op, e.Path, e.Err)
	}
}

// A package can also define its own specific Error with additional methods, like net.Error:
// https://github.com/golang/go/blob/master/src/net/net.go#L207

//type Error interface {
//	error
//	Timeout() bool   // Is the error a timeout?
//	Temporary() bool // Is the error temporary?
//}

// Client code can test for a net.Error with a type assertion and then
// distinguish transient network errors from permanent ones.
func dialTCP() {
	_, err := net.DialTimeout("tcp", "google1.com:80", time.Duration(2*time.Nanosecond))
	if err != nil {
		if e, ok := err.(net.Error); ok {
			if e.Timeout() {
				fmt.Println("request timedout")
			}

			if e.Temporary() {
				// retry logic
			}
		}
	}
}

func main() {
	readFile()
	dialTCP()
}
