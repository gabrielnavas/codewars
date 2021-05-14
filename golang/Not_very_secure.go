package main

import (
	"fmt"
)

func alphanumeric(str string) bool {
	if str == "" {
		return false
	}
	for _, chr := range str {
		if int(chr) < int('A') || int(chr) > int('z') {
			if int(chr) < int('0') || int(chr) > int('9') {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(alphanumeric(".*?") == false)
	fmt.Println(alphanumeric("a") == true)
	fmt.Println(alphanumeric("Mazinkaiser") == true)
	fmt.Println(alphanumeric("hello world_") == false)
	fmt.Println(alphanumeric("PassW0rd") == true)
	fmt.Println(alphanumeric("     ") == false)
	fmt.Println(alphanumeric("") == false)
	fmt.Println(alphanumeric("\n\t\n") == false)
	fmt.Println(alphanumeric("ciao\n$$_") == false)
	fmt.Println(alphanumeric("__ * __") == false)
	fmt.Println(alphanumeric("&)))(((") == false)
	fmt.Println(alphanumeric("43534h56jmTHHF3k") == true)
}
