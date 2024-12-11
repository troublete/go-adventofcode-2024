package main

import "fmt"

func main() {
	p := ContentToPuzzle(input)
	xs := p.FindAll('X')
	total := 0
	for _, x := range xs {
		n, _ := x.CheckFor("XMAS")
		total += n
	}
	fmt.Printf("total xmas': %v\n", total)
}
