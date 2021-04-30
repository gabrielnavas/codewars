package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	s = strings.ToLower(s)
	splited := strings.Split(s, "")
	strFinal := ""
	for i, chr := range splited {
		count := i + 1
		strFinal += strings.ToUpper(chr)
		count--
		for count > 0 {
			strFinal += chr
			count--
		}
		strFinal += "-"
	}
	return strFinal[:len(strFinal)-1]
}

func main() {
	result := Accum("ZpglnRxqenU")
	testcase := "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu"
	if result == testcase {
		fmt.Println(result)
	}
}
