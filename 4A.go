package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a == 2 {
		fmt.Println("NO")
		return
	}
	m := a % 2
	if m == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
