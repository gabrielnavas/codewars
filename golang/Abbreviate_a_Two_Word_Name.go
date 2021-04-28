package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	splited := strings.Split(name, " ")
	firstChr := strings.ToUpper(string(splited[0][0]))[0]
	secondChr := strings.ToUpper(string(splited[1][0]))[0]
	abbreviateName := fmt.Sprintf("%c.%c", firstChr, secondChr)
	return abbreviateName
}

func main() {
	fmt.Println(AbbrevName("Gabriel Navas"))
}
