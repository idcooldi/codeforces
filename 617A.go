package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	var count, step int
	for i := 5; i > 0; i-- {
		step = x / i
		x = x - step*i
		count += step
		if x <= 0 {
			break
		}
	}
	fmt.Println(count)
}
