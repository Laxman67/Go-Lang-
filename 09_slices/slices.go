package main

import "fmt"

func main() {

	// Q1. Ek slice banao marks := []int{85, 90, 78, 92, 70}
	marks := []int{85, 1, 78, 5, 70}
	// var sum int = 0
	// for i := range marks {
	// 	sum = sum + marks[i]
	// 	fmt.Println(marks[i])
	// }
	// fmt.Println("Total: ", sum)

	//Q2 - Print karo: total marks (loop se sum karke), highest, lowest
	// students := []string{"Rahul", "Priya", "Amit", "Sneha"}
	// fmt.Println(students[1:3])
	// students = append(students, "vikas")
	// fmt.Println(students)

	// Q3. (Tricky ⚠️)
	// original := []int{1, 2, 3}
	// copySlice := original
	// copySlice[0] = 99
	// fmt.Println(original)
	res := filterPass(marks)
	fmt.Println("Result: ", res)

}
func filterPass(marks []int) []int {
	var passing []int
	for _, m := range marks {
		if m > 60 {
			passing = append(passing, m)
		}

	}
	return passing

}
