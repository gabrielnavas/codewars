package main

import "fmt"

func ProductFib(param uint64) [3]uint64 {
	var (
		f1 uint64 = 1
		f2 uint64 = 1
	)
	for {
		if f1*f2 >= param {
			break
		}
		f1, f2 = f2, f2+f1
	}
	if f1*f2 == param {
		return [3]uint64{f1, f2, 1}
	}
	return [3]uint64{f1, f2, 0}
}

func main() {
	fmt.Println(ProductFib(4895) == [3]uint64{55, 89, 1})
	fmt.Println(ProductFib(5895) == [3]uint64{89, 144, 0})
	fmt.Println(ProductFib(74049690) == [3]uint64{6765, 10946, 1})
	fmt.Println(ProductFib(84049690) == [3]uint64{10946, 17711, 0})
}
