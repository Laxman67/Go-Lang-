package main

import (
	"fmt"
	"strings"
)

func main() {

	// Task 3:
	text := "go is great and go is fast"
	mapper := make(map[string]int)

	for _, word := range strings.Fields(text) {
		// fmt.Println(_) //key
		mapper[word]++ // count badhao
	}

	fmt.Println(mapper)

	// Taks 2:
	// scores := map[string]int{"Rahul": 85, "Priya": 92, "Amit": 78, "Sneha": 90}

	// if initialise; value {} else{}

	// if v, ok := scores["Rahul"]; ok {
	// 	fmt.Println("Found!", v)
	// } else {
	// 	fmt.Println("Found!", v)

	// }

	// Taks 1:

	// var sum, highest int

	// for key, value := range scores {
	// 	if value > highest {
	// 		highest = value
	// 	}
	// 	sum = sum + value
	// 	fmt.Printf("%s => %d\n", key, value)

	// }

	// fmt.Println("Total Marks: ", sum)
	// fmt.Println("Highest Marks :", highest)

	// // student:=map[keyType]valueType
	// capitals := make(map[string]string) ///decalration
	// capitals["IN"] = "Delhi"
	// capitals["CH"] = "Beijing"

	// capitals["US"] = "Washington"

	// // Check if Present

	// val, ok := capitals["US"]
	// fmt.Println(capitals)

	// fmt.Println("ok", ok)
	// fmt.Println("val:", val)

	// fmt.Println("Looping....")
	// for key, val := range capitals {
	// 	fmt.Printf("%s => %s\n", key, val)
	// }

	// fmt.Println("Length: ", len(capitals))

	// m := map[string]int{"Rohit": 12, "Laksh": 45}

	// if v, ok := m["Rohit"]; ok {
	// 	fmt.Println("Found !", v)
	// } else {
	// 	fmt.Println("Not Found!")
	// }
}
