package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	var upper, lower int
	l := func(c rune) bool {
		return unicode.IsLower(c)
	}
	u := func(c rune) bool {
		return unicode.IsUpper(c)
	}
	for _, v := range r {
		if u(v) {
			upper++
		}
		if l(v) {
			lower++
		}
	}
	if lower > upper || lower == upper {
		fmt.Println(strings.ToLower(s))
		return
	}
	fmt.Println(strings.ToUpper(s))
}
