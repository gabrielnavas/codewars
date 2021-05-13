package main

import (
	"fmt"
)

func Solution(word string) string {
	chrs := []rune(word)
	for i, j := 0, len(chrs)-1; i < j; i, j = i+1, j-1 {
		chrs[i], chrs[j] = chrs[j], chrs[i]
	}
	return string(chrs)
}

func main() {
	fmt.Println(Solution("world") == "dlrow")
}
