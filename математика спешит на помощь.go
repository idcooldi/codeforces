package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var a, c string
	fmt.Scan(&a)
	b := strings.Split(a, "+")
	sort.Sort(sort.StringSlice(b))
	for _, v := range b {
		c += v + "+"
	}
	fmt.Println(strings.TrimRight(c, "+"))
}
