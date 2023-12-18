package main

import "fmt"

func main() {
	var length int
	fmt.Scanf("%d\n", &length)

	array := make([]string, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%s\n", &array[i])
	}
	var count int
	for _, v := range array {
		if v == "X++" || v == "++X" {
			count++
		}
		if v == "X--" || v == "--X" {
			count--
		}
	}
	fmt.Println(count)
}
