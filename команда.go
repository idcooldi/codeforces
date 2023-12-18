package main

import (
	"fmt"
)

func main() {
	var length int
	fmt.Scanf("%d\n", &length)
	var sum int
	for i := 0; i < length; i++ {
		var x, j, k int
		fmt.Scanf("%d %d %d\n", &x, &j, &k)
		if x+j+k > 1 {
			sum++
		}
	}

	fmt.Println(sum)
}
