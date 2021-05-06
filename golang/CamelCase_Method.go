package main

import (
	"fmt"
	"strings"
)

func CamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	var (
		phrase  string
		newWord string
	)
	words := strings.Split(s, " ")
	for _, word := range words {
		if len(word) > 0 {
			newWord = strings.ToUpper(string(word[0])) + word[1:]
			phrase += newWord
		}
	}
	return phrase
}

func main() {
	t := [...][2]string{
		{"test case", "TestCase"},
		{"camel case method", "CamelCaseMethod"},
		{"say hello ", "SayHello"},
		{" camel case word", "CamelCaseWord"},
		{"", ""},
	}

	for _, v := range t {
		fmt.Printf("\"%s\" == \"%v\" => \"%v\"\n", v[0], v[1], CamelCase(v[0]) == v[1])
	}
}
