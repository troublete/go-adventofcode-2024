package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	calcs := Extract(input)
	sum := 0.0
	for _, c := range calcs {
		sum += c.Mul()
	}

	secondSum := 0.0
	for _, s := range Sections(input) {
		for _, e := range Extract(s) {
			secondSum += e.Mul()
		}
	}

	fmt.Printf("sum: %.f\n", sum)
	fmt.Printf("sectioned sum: %.f\n", secondSum)
}
