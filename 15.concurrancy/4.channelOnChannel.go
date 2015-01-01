package main

import (
	"fmt"
	"time"
)

// It is possible to send other channels over a channel. The chChan parameter of
// emmit function below is used to sends another channel back to caller.
func emmit(done <-chan struct{}, chChan chan<- chan string, s []string) {
	i, l := 0, len(s)
	wchan := make(chan string)
	chChan <- wchan

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
	chChan := make(chan chan string)

	i := 0

	strs := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	go emmit(done, chChan, strs)

	wordChan := <-chChan

	for word := range wordChan {
		fmt.Printf("%s ", word)

		//if i > 15 {
		//	// send a value down the done channel
		//	done <- struct{}{}
		//}

		i += 1
	}
}
