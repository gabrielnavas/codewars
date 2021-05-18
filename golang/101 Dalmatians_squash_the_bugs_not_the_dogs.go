package main

import "fmt"

func HowManyDalmatians(number int) string {
	dogs := []string{
		"Hardly any",
		"More than a handful!",
		"Woah that's a lot of dogs!",
		"101 DALMATIONS!!!",
	}
	if number <= 10 {
		return dogs[0]
	}
	if number <= 50 {
		return dogs[1]
	}
	if number < 101 {
		return dogs[2]
	}
	return dogs[3]
}

func main() {
	fmt.Println(HowManyDalmatians(26) == "More than a handful!")
	fmt.Println(HowManyDalmatians(8) == "Hardly any")
	fmt.Println(HowManyDalmatians(14) == "More than a handful!")
	fmt.Println(HowManyDalmatians(80) == "Woah that's a lot of dogs!")
	fmt.Println(HowManyDalmatians(100) == "Woah that's a lot of dogs!")
	fmt.Println(HowManyDalmatians(50) == "More than a handful!")
	fmt.Println(HowManyDalmatians(10) == "Hardly any")
	fmt.Println(HowManyDalmatians(101) == "101 DALMATIONS!!!")
}
