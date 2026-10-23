package main

import "fmt"

// basic syntax
func greet(name string) string {
	return "Hello, " + name
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("%d Cannot divide by zero", a)
	}
	return a / b, nil

}
func sum(nums ...int) int { // ... = slice ban jaata hai
	total := 0
	// key _ , value:=range nums
	for _, n := range nums {
		total += n
	}
	return total
}

// Pass by value and Pointer

// Closure
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func change(n *int) {
	*n = 99
}

func main() {
	a := 90
	fmt.Println("Original: ", a)
	change(&a)
	fmt.Println("Original :After: ", a)

	// c := counter()
	// fmt.Println(c())
	// fmt.Println(c())
	// fmt.Println(c())

	// sumResult := sum(1, 4, 1, 5, 6, 90)
	// println("Sum Result is : ", sumResult)

	// sum(1, 5, 1, 6, 77, 90)
	// result, err := divide(3, 89)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("result", result)

}
