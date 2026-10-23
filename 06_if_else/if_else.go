package main

import "fmt"

func main() {
	// if age >= 18 {
	// 	fmt.Println("Is Adult")
	// } else {
	// 	fmt.Println("Not an  Adult")

	// }

	if age := 122; age >= 18 {
		fmt.Println("Is Adult")
	} else if age >= 12 {
		fmt.Println("Person in Teenager")

	} else {
		fmt.Println("Person in Kid")

	}

}
