package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var in []byte

func main() {
	m := ContentToMap(in)
	p := m.FindAllPaths(0, 9)

	fmt.Printf("score: %v\n", len(p))
}
