package main

import (
	"fmt"
)

func main() {
	var length, win int
	fmt.Scanf("%d %d\n", &length, &win)
	array := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &array[i])

	}

	var count int
	for _, v := range array {
		if array[win-1] <= v && v != 0 {
			count++
		}
	}
	fmt.Println(count)
}
