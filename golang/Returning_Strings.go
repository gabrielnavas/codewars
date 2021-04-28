package main

import "fmt"

func Greet(name string) string {
	str := fmt.Sprintf("Hello, %s how are you doing today?", name)
	return str
}

func main() {
	fmt.Println(Greet("gabriel"))
}
