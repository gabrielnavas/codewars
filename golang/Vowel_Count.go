package main

import "fmt"

func Any(options []rune, test_chr rune) bool {
	for _, chr := range options {
		if chr == test_chr {
			return true
		}
	}
	return false
}

func GetCount(str string) (count int) {
	for _, chr := range str {
		if Any([]rune{'a', 'e', 'i', 'o', 'u'}, chr) {
			count++
		}
	}
	return
}

func main() {
	fmt.Println(GetCount("abracadabra") == 5)
}
