package main

import (
	"fmt"
	"time"
)

func ProcessNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Received:", num)
	}
}

func sum(result chan int, num1, num2 int) {
	sum := num1 + num2
	result <- sum

}

func task(done chan bool) {
	defer func() {
		done <- true
	}()

	fmt.Println("Processings")
}

func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true
	}()
	for email := range emailChan {
		fmt.Println("Sending Email to: ", email)
		time.Sleep(time.Second)
	}
}

func main() {
	emailChan := make(chan string, 5)
	done := make(chan bool)
	go emailSender(emailChan, done)
	for i := 0; i < 6; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	close(emailChan)
	<-done

	fmt.Println("Done Email Sending...")

	// done := make(chan bool)
	// go task(done)
	// fmt.Println("he;p")
	// result := make(chan int)
	// go sum(result, 12, 46)

	// res := <-result

	// fmt.Println(res)
	// numChan := make(chan int)

	// go ProcessNum(numChan)

	// for {
	// 	numChan <- rand.Intn(20)
	// }
}
