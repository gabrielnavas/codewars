package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DigitalRoot(n int) int {
	var (
		r   int
		sum int
	)
	for n > 9 {
		sum = 0
		for _, numChr := range strings.Split(strconv.Itoa(n), "") {
			r, _ = strconv.Atoi(numChr)
			sum += r
		}
		n = sum
	}
	return n
}

func main() {
	fmt.Println(DigitalRoot(16) == 7)
	fmt.Println(DigitalRoot(195) == 6)
	fmt.Println(DigitalRoot(992) == 2)
	fmt.Println(DigitalRoot(167346) == 9)
	fmt.Println(DigitalRoot(0) == 0)
}
