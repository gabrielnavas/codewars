package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
	splited := strings.Split(str, " ")
	for i, word := range splited {
		if len(word) >= 5 {
			splited[i] = reverseStr(word)
		}
	}
	return strings.Join(splited, " ")
}

func reverseStr(str string) string {
	var reversedStr []rune
	for i := len(str) - 1; i > -1; i-- {
		reversedStr = append(reversedStr, rune(str[i]))
	}
	return string(reversedStr)
}

func main() {
	fmt.Println(SpinWords("Welcome") == "emocleW")
	fmt.Println(SpinWords("to") == "to")
	fmt.Println(SpinWords("CodeWars") == "sraWedoC")
	fmt.Println(SpinWords("Hey fellow warriors") == "Hey wollef sroirraw")
	fmt.Println(SpinWords("Burgers are my favorite fruit") == "sregruB are my etirovaf tiurf")
	fmt.Println(SpinWords("Pizza is the best vegetable") == "azziP is the best elbategev")

}
