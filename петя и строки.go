package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, y string
	fmt.Scanf("%s\n %s", &x, &y)

	r1 := []rune(strings.ToLower(x))
	r2 := []rune(strings.ToLower(y))

	for i := 0; i < len(r1); i++ {
		if r1[i] > r2[i] {
			fmt.Println(1)
			return
		}
		if r1[i] < r2[i] {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(0)
}
