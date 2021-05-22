package main

import (
	"fmt"
	"strconv"
)

func PrinterError(s string) string {
	errors := 0
	for _, chr := range s {
		if chr > rune('m') {
			errors++
		}
	}
	return strconv.Itoa(errors) + "/" + strconv.Itoa(len(s))
}

func main() {
	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz") == "3/56")
	fmt.Println(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz") == "6/60")
	fmt.Println(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu") == "11/65")
}
