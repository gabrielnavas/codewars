package main

import (
	"fmt"
	"strings"
)

func toWeirdCase(str string) string {
	var strReturn string
	for _, word := range strings.Split(str, " ") {
		for i, chr := range strings.Split(word, "") {
			if i%2 == 0 {
				strReturn += strings.ToUpper(chr)
			} else {
				strReturn += strings.ToLower(chr)
			}
		}
		strReturn += " "
	}
	lessLastSpace := strReturn[:len(strReturn)-1]
	return lessLastSpace
}

func main() {

	fmt.Println(toWeirdCase("abc def") == "AbC DeF")
	fmt.Println(toWeirdCase("ABC") == "AbC")
	fmt.Println(toWeirdCase("This is a test Looks like you passed") == "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd")
}
