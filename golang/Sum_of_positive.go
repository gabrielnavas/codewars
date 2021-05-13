package main

import "fmt"

func PositiveSum(numbers []int) (sum int) {
	for _, n := range numbers {
		if n > 0 {
			sum += n
		}
	}
	return
}

func main() {
	fmt.Println(PositiveSum([]int{1, 2, 3, 4, 5}) == 15)
	fmt.Println(PositiveSum([]int{1, -2, 3, 4, 5}) == 13)
	fmt.Println(PositiveSum([]int{}) == 0)
	fmt.Println(PositiveSum([]int{-1, -2, -3, -4, -5}) == 0)
	fmt.Println(PositiveSum([]int{-1, 2, 3, 4, -5}) == 9)
}
