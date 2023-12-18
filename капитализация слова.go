package main

import (
	"fmt"
	"unicode"
)

func main() {
	var a string
	fmt.Scan(&a)
	r := []rune(a)
	r[0] = unicode.ToUpper(r[0])
	fmt.Println(string(r))
}
