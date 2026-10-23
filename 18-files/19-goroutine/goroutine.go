package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	go func() {
		time.Sleep(200 * time.Microsecond)
		ch3 <- "Response from service C"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Response from service A"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch4 <- "Response from service D"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "Response from service B"
	}()

	// Whichever channel responds first gets picked
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case msg3 := <-ch3:
		fmt.Println(msg3)
	case msg4 := <-ch3:
		fmt.Println(msg4)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}
}
