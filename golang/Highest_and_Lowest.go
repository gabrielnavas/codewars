package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func toArrayInt(numsStr []string) []int {
	nums := []int{}
	for i := 0; i < len(numsStr); i++ {
		chr, _ := strconv.Atoi(numsStr[i])
		nums = append(nums, chr)
	}
	return nums
}

func mountFinalStr(nums []int) string {
	lastNum := nums[len(nums)-1]
	firstNum := nums[0]
	finalStr := fmt.Sprintf("%d %d", lastNum, firstNum)
	return finalStr
}

func HighAndLow(in string) string {
	splited := strings.Split(in, " ")
	nums := toArrayInt(splited)
	sort.Ints(nums)
	finalStr := mountFinalStr(nums)
	return finalStr
}

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}
