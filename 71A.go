package main

import (
	"fmt"
	"os"
)

func main() {
	var length int
	fmt.Fscan(os.Stdin, &length)

	for i := 0; i < length; i++ {
		var text string
		fmt.Fscan(os.Stdin, &text)
		fmt.Println(cut(text, 10))
	}
}
func cut(text string, i int) string {
	l := len(text)
	if l <= i {
		return text
	}
	runes := []rune(text)
	start := string(runes[:1])
	end := string(runes[l-1:])

	return fmt.Sprintf("%s%d%s", start, l-2, end)
}
