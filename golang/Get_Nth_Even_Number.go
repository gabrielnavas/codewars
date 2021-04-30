package main

import "fmt"

func NthEven(n int) int {
	nTwo := n * 2
	nTwo--
	for nTwo%2 != 0 {
		nTwo--
	}
	return nTwo
}

func main() {
	fmt.Println(NthEven(1) == 0)
	fmt.Println(NthEven(2) == 2)
	fmt.Println(NthEven(3) == 4)
	fmt.Println(NthEven(1298734) == 2597466)
}
