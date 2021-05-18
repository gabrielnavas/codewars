package main

import "fmt"

func Grow(arr []int) (sum int) {
	sum = arr[0]
	for _, n := range arr[1:] {
		sum *= n
	}
	return
}

func main() {
	fmt.Println(Grow([]int{1, 2, 3}) == 6)
	fmt.Println(Grow([]int{4, 1, 1, 1, 4}) == 16)
	fmt.Println(Grow([]int{2, 2, 2, 2, 2, 2}) == 64)
}
