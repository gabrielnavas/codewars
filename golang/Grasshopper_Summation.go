package main

import "fmt"

func Summation(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

func main() {
	examples := [...][2]int{
		{1, 1},
		{8, 36},
		{22, 253},
		{100, 5050},
		{213, 22791},
	}
	for i := 0; i < len(examples); i++ {
		fmt.Println(Summation(examples[i][0]) == examples[i][1])
	}
}
