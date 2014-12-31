package main

import "fmt"

// A channel provides a mechanism for concurrently executing functions (goroutinues) to communicate by
// sending and receiving values of a specified element type, i.e. channels are typed, you can only send
// or receive values of the specified type.

// The optional <- operator specifies the channel direction, send or receive. If no direction is given,
// the channel is bidirectional. A channel may be constrained only to send or only to receive by
// conversion or assignment.

// A new, initialized channel value can be made using the built-in function make, which takes the
// channel type and an optional capacity as arguments. The capacity, in number of elements, sets
// the size of the buffer in the channel. If the capacity is zero or absent, the channel is
// unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise,
// the channel is buffered and communication succeeds without blocking if the buffer is not full (sends)
// or not empty (receives). A nil channel is never ready for communication.

// A channel may be closed with the built-in function close. The multi-valued assignment form of the
// receive operator reports whether a received value was sent before the channel was closed. A closed
// channel never blocks, and would return the zero value of the channel type. It is idiomatic for only
// for a sender to close a channel.

// Channel is also synchronization operation, for unbuffered channel a send cannot proceed until someone
// is ready to receive on the same channel. Likewise buffered channel will block and behave like unbuffered
// channel when the channel capacity is reached.

// see these excellent blog posts for more on this subject
// http://dave.cheney.net/2013/04/30/curious-channels
// http://guzalexander.com/2013/12/06/golang-channels-tutorial.html
// http://blog.golang.org/pipelines
// http://blog.golang.org/context

// emmit sends strings over the channel
func emmit(c chan string) {
	strs := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	for _, v := range strs {
		//send v into the channel c. The arrow indicate the direction of data flow.
		c <- v
	}

	// Closing is only necessary when the receiver must be told there are no more values coming,
	// such as to terminate a range loop. Since we are looping over the out channel in the main
	// function, failure to see close this channel would result in deadlock. Uncomment the following
	// line to see the deadlock in action
	close(c)
}

func main() {

	// initialize a channnel of string
	out := make(chan string)

	// execute the emmit function in separate goroutinue
	go emmit(out)

	// receive a value from a channel
	first := <-out
	fmt.Printf("%s\n", first)

	// iterate over the channel.
	// If the channel being iterated over is not closed we will run into deadlock
	for v := range out {
		fmt.Printf("%s ", v)
	}

}
