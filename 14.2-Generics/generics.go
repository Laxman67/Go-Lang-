package main

import (
	"fmt"
)

func printSlice[T comparable](items []T) {
	for _, item := range items {

		fmt.Println(item)
	}
}
func main() {
	printSlice([]int{1, 65, 24, 1, 23})
	printSlice([]string{"Hello", "World"})
	printSlice([]bool{true})

}
