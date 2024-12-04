package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	a, b := inputToLists(input)
	d := distance(a, b)
	s := similarity(a, b)

	fmt.Printf("distance: %.0f\n", d)
	fmt.Printf("similarity: %.0f\n", s)
}
