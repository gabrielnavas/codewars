package main

import (
	"fmt"
	"math"
)

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	var difference = dadYearsOld - (sonYearsOld * 2)
	return int(math.Abs(float64(difference)))
}

func main() {
	fmt.Println(TwiceAsOld(36, 7))
}
