package main

import "fmt"

func main() {
	var price, balance, amount, sum int
	fmt.Scan(&price, &balance, &amount)

	for i := amount; i > 0; i-- {
		sum += price * i
	}
	credit := sum - balance
	if credit < 1 {
		fmt.Println(0)
		return
	}
	fmt.Println(credit)
}
