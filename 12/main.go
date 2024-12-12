package main

import (
	"fmt"
)

func main() {
	m := ContentToMap(in)
	regions := m.FindRegions()
	total := 0
	for _, r := range regions {
		total += r.FencePrice()
	}
	totalBulk := 0
	for _, r := range regions {
		totalBulk += r.FencePriceBulk()
	}
	fmt.Printf("fence price: %v\n", total)
	//fmt.Printf("fence price bulk: %v\n", totalBulk)
}
