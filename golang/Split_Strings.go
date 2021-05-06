package main

import "fmt"

func Solution(str string) []string {
	var resp []string
	var twoByte [2]byte
	lenStr := len(str)
	for i := 0; i < lenStr; {
		for j := 0; i < lenStr && j < 2; j++ {
			twoByte[j] = byte(str[i])
			i++
		}
		resp = append(resp, string(twoByte[:]))
	}
	if lenStr%2 != 0 {
		lenLastItem := len(resp) - 1
		lastStr := resp[lenLastItem]
		resp[lenLastItem] = string(lastStr[0]) + string("_")
	}
	return resp
}

func main() {
	fmt.Println(Solution("abc"))    // == []string{"ab", "c_"})
	fmt.Println(Solution("dawsd"))  // == []string{"da", "ws", "d_"})
	fmt.Println(Solution("awsaws")) //  == []string{"aw", "sa", "ws"})
}
