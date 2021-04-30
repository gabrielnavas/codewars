package main

import (
	"fmt"
	"strings"
)

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

func main() {
	fmt.Println(DNAtoRNA("GCAT") == "GCAU")
}
