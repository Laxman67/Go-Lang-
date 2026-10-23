package main

import "fmt"

func fullName(fname, lname string) string {
	return fname + " " + lname

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil

}
func main() {
	// fmt.Println("Hello : ", fullName("Laxman", "NA"))
	result, err := divide(34, 9)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(result)

}
