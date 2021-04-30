package main

import (
	"fmt"
	"strings"
)

func EncryptThis(text string) string {
	if len(text) == 0 {
		return text
	}
	splited := strings.Split(text, " ")
	finalStr := ""
	var wordCrypt string = ""
	for _, word := range splited {
		if len(word) == 1 {
			wordCrypt = fmt.Sprintf("%d", word[0])
		} else if len(word) == 2 {
			wordCrypt = fmt.Sprintf("%d%c", word[0], word[1])
		} else if len(word) == 3 {
			wordCrypt = fmt.Sprintf("%d%c%c",
				word[0],
				word[len(word)-1],
				word[1],
			)
		} else {
			wordCrypt = fmt.Sprintf("%d%c%s%c",
				word[0],
				word[len(word)-1],
				word[2:len(word)-1],
				word[1],
			)
		}
		finalStr += fmt.Sprintf("%s ", wordCrypt)
	}
	lessSpace := finalStr[:len(finalStr)-1]
	return lessSpace
}

func main() {
	fmt.Println(EncryptThis("") == "")
	fmt.Println(EncryptThis("A wise old owl lived in an oak") == "65 119esi 111dl 111lw 108dvei 105n 97n 111ka")
	fmt.Println(EncryptThis("A wise old owl lived in an oak"))
	// fmt.Println(EncryptThis("The more he saw the less he spoke") == "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp")
	// fmt.Println(EncryptThis("The less he spoke the more he heard") == "84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare")
	// fmt.Println(EncryptThis("Why can we not all be like that wise old bird") == "87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri")
	// fmt.Println(EncryptThis("Thank you Piotr for all your help") == "84kanh 121uo 80roti 102ro 97ll 121ruo 104ple")
}
