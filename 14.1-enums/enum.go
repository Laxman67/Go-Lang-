package main

import "fmt"

// enumrated types
type OrderStatus string

const (
	Recieved  OrderStatus = "recieved"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing Order status to: ", status)
}
func main() {
	changeOrderStatus(Recieved)
	changeOrderStatus(Confirmed)
	changeOrderStatus(Prepared)

}
