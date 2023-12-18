package main

import "fmt"

func main() {
	var l, count int
	var a string
	fmt.Scan(&l)
	fmt.Scan(&a)
	r := []rune(a)
	if l < 2 {
		fmt.Println(0)
		return
	}
	for k, v := range r {
		if k == l-1 {
			break
		}
		if v == r[k+1] {
			count++
		}
	}
	if count == l {
		count--
	}
	fmt.Println(count)
}
