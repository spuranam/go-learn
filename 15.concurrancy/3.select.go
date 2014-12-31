package main

import (
	"fmt"
	"time"
)

// A "select" statement chooses which of a set of possible send or receive operations will proceed.
// It looks similar to a "switch" statement but with the cases all referring to communication operations.

// Execution of a "select" statement proceeds in several steps:

// For all the cases in the statement, the channel operands of receive operations and the channel and right-hand-side
// expressions of send statements are evaluated exactly once, in source order, upon entering the "select" statement.
// The result is a set of channels to receive from or send to, and the corresponding values to send. Any side effects
// in that evaluation will occur irrespective of which (if any) communication operation is selected to proceed.
// Expressions on the left-hand side of a RecvStmt with a short variable declaration or assignment are not yet evaluated.

// If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random
// selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement
// blocks until at least one of the communications can proceed.

// Unless the selected case is the default case, the respective communication operation is executed.

// If the selected case is a RecvStmt with a short variable declaration or an assignment, the left-hand side expressions
// are evaluated and the received value (or values) are assigned.

// The statement list of the selected case is executed.

func emmit(done <-chan struct{}, wchan chan<- string, s []string) {
	i, l := 0, len(s)

	defer close(wchan)

	t := time.NewTimer(2 * time.Nanosecond)
	//t := time.After(1 * time.Nanosecond)

	for {

		select {
		// we are ready to deliver a value on the wchan channel
		case wchan <- s[i]:
			i += 1
			if i >= l {
				i = 0
			}
		// we recieved a value on a done channel, hence abort
		case <-done:
			fmt.Println("Recieved a done signal")
			return
		// timeout has been breached, hence abort
		case <-t.C:
			fmt.Println("timed out")
			return
		}

	}

}

func main() {

	done := make(chan struct{})
	wordChan := make(chan string)

	i := 0

	strs := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	go emmit(done, wordChan, strs)

	for word := range wordChan {
		fmt.Printf("%s ", word)

		//if i > 15 {
		//	// send a value down the done channel
		//	done <- struct{}{}
		//}

		i += 1
	}
}
