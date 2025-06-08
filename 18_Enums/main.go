package main

import "fmt"

// Enumerated Types ==========================
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Delivered
)

//=============================================
func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to ====>", status)
}

func main() {
	changeOrderStatus(Received)
}
