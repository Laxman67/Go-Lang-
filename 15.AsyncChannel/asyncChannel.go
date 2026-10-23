package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d Started \n", i)

	fmt.Printf("Worker %d End \n", i)
	wg.Done()

}
func main() {

	fmt.Println("Explore goroutine started")

	var wg sync.WaitGroup
	// start 3 workser goroutine
	for i := 1; i <= 3; i++ {
		wg.Add(1) //increement the waitgroup Counter
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Workers taks Completed")
}
