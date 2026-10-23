package main

import "fmt"

func main() {

	whoAmi := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer", t)
		case string:
			fmt.Println("String")

		case bool:
			fmt.Println("Boolean")

		default:
			fmt.Println("Other")
		}

	}
	var name = "laksh"
	whoAmi(name)

	// multiple condition
	// switch time.Now().Weekday() {

	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's Weekend")
	// default:
	// 	fmt.Println("It's  Workday! sheet")
	// }

	// Simple Switch
	// i := 111
	// switch i {
	// case 1:
	// 	fmt.Println(1)
	// case 2:
	// 	fmt.Println(2)
	// default:
	// 	fmt.Println("Others")
	// }

}
