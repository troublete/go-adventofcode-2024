package main

import "fmt"

func main() {
	l := ContentToLine(input)
	for c := 0; c < 25; c++ {
		fmt.Println(c)
		l = ApplyRules(l)
	}
	fmt.Printf("n stones: %v", len(l))
}
