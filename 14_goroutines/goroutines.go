package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello World")
	time.Sleep(2 * time.Second)
	fmt.Println("SayHello endded after 2 seconds")
}

func sayHello2() {
	fmt.Println("Hello World: 2")
	time.Sleep(2 * time.Second)
	fmt.Println("SayHello:2  endded after 2 seconds")
}

func sayHi() {
	fmt.Println("Hi ")
}
func main() {
	fmt.Println("Learning goroutine")
	go sayHello()
	go sayHello2()
	time.Sleep(2 * time.Second)
	sayHi()

}
