package main

import "fmt"

func main() {
	v := ContentToValidator(in)
	total := 0
	for _, vm := range v.ValidManuals() {
		total += vm.MiddlePage()
	}
	fmt.Printf("middle page sum: %v\n", total)
}
