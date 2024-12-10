package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	m := MapFromInput(input)
	l := m.Layout().Improve(nil)
	fmt.Printf("line: %v\n", l.ToString())
	fmt.Printf("checksum: %v\n", l.Checksum())
}
