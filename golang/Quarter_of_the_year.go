package main

import "fmt"

func QuarterOf(month int) int {
	if month <= 3 {
		return 1
	}
	if month <= 6 {
		return 2
	}
	if month <= 9 {
		return 3
	}
	return 4
}

func main() {
	fmt.Println(QuarterOf(3))
	fmt.Println(QuarterOf(8))
	fmt.Println(QuarterOf(11))
}
