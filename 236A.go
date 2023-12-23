package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	r := []rune(a)
	for {
		for k, v := range r {
			i := strings.Count(string(r), string(v))
			if i > 1 {
				r = delete(r, k, k+1)
				break
			}
			if k == len(r)-1 {
				goto END
			}
		}
	}
END:
	if len(r)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
		return
	}
	fmt.Println("IGNORE HIM!")
}

func delete[S ~[]E, E any](s S, i, j int) S {
	_ = s[i:j] // bounds check

	return append(s[:i], s[j:]...)
}
