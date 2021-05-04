package main

import "fmt"

func FindUniq(arr []float32) float32 {
	nums := map[float32]float32{}
	var unique float32 = arr[0]

	for _, n := range arr {
		if nums[n] == 0 {
			nums[n] = 1
		} else {
			nums[n]++
		}
	}
	for key, value := range nums {
		if value == 1 {
			unique = key
		}
	}
	return unique
}

func main() {
	fmt.Println(FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}) == 2)
}
