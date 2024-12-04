package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	in := inputToReports(input)
	fmt.Printf("save reports: %v", safetyReport(in))
}
