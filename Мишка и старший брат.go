package main

import (
	"fmt"
)

func main() {
	var count int
	var a, b float64
	fmt.Scan(&a, &b)
	for {
		count++
		a *= 3
		b *= 2
		if a > b {
			fmt.Println(count)
			break
		}
	}
}
