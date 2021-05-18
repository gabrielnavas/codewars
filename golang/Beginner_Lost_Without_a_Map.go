package main

func Maps(x []int) []int {
	doubles := []int{}
	for len, i := len(x), 0; i < len; i++ {
		doubles = append(doubles, x[i]*2)
	}
	return doubles
}

func main() {

}
