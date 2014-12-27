package main

import (
	"fmt"
	"time"
)

func greet(s string) {
	fmt.Println(s)
}

func main() {
	// A goroutinue is an independent and concurrently executing computation.

	// Some intresting properties of a goroutinue are:
	// 1. They are extremly light weight and cheap when compared to threads.
	// 2. Hence you can create hundreds even millions on a commodity hardware.
	// 3. All goroutinues spwanned by an application runs in the same address space
	// 4. They are mutiplexed on to the OS threads by Go runtime scheduler.
	// 5. They run independently of other goroutines, in the same address space.
	// 6. When the main goroutinue exits the program terminates.
	// 7. Hence you must arrange for the main goroutinue to hang around until all
	// other gorountinues either terminate or are cancelled.
	// 8. All Go programs have atleast one goroutinue (main goroutinue)
	// associated with it. The main function runs in its own goroutinue.

	// A "go" statement starts the execution of a function call as an
	// independent concurrent thread of control, or goroutine, within the
	// same address space. The passes function, and its arguments are
	// evaluated in the context of current goroutinue and the passed function
	// itself is executed in a new separete goroutinue.
	go greet("there!")
	fmt.Println("Hi")

	// Since the program terminates when the main goroutinue returns, if we
	// to remove the cal to sleep  we will never see the output of the function
	// greet, since it is executing in a separate goroutine
	time.Sleep(10 * time.Nanosecond)
}
