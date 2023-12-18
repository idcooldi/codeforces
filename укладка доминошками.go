package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanf("%d %d\n", &x, &y)

	z := x * y
	a := z / 2
	var b int = int(a)
	fmt.Println(b)
}
