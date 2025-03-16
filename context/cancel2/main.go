package main

import (
	"fmt"
	"time"
)


func main() {
	chan1 := make(chan bool, 2)
	chan2 := make(chan bool, 2)

	go testFunc(chan1, chan2, "call-1")
	go testFunc(chan1, chan2, "call-2")

	time.Sleep(100 * time.Millisecond)

	chan2 <- true
	chan2 <- true

	<-chan2
	<-chan2

	fmt.Println("Back to main go routine !!")

}


func testFunc(ch chan bool, ch2 chan bool, msg string) <- chan bool {
	fmt.Printf("Reached to `testFunc` with message: %s\n", msg)

	select {
	case <- time.After(500 * time.Millisecond):
		fmt.Printf("Executing with message: %s\n", msg)
	case ch2 <- <-ch:
		fmt.Printf("Cancelling function with message: %s\n", msg)
		return ch2
	}

	fmt.Printf("Finishing function with message: %s\n", msg)
	ch2 <- true
	return ch2

}