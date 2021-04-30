package main

import (
	"fmt"
)

func CorrectTail(body string, tail rune) bool {
	if rune(body[len(body)-1]) == tail {
		return true
	}
	return false
}

func main() {
	fmt.Println(CorrectTail("Fox", 'x') == true)
}
