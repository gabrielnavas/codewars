package main

import (
	"fmt"
)

func RemoveChar(word string) string {
	lessFirst := word[1:]
	lenWordLessOne := len(lessFirst) - 1
	lessLast := lessFirst[:lenWordLessOne]
	return lessLast
}

func main() {
	fmt.Println(RemoveChar("Gabriel"))
}
