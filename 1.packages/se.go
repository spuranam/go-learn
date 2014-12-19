package main

// import _ "./sideeffect" - Suppresses compiler warnings related to
// sideeffect if it is not being used, and executes initialization functions
// if there are any. The remainder of sideeffect is inaccessible.
// This form of imports are generally used to for the sideeffect, for example
// when dealing with database drivers you would import the drivers using this.
// for details see. https://code.google.com/p/odbc/

// in this example importing package sideeffect, execute its init function
// which casuse it to print a message to stdout.

import _ "./sideeffect"

func main() {

}
