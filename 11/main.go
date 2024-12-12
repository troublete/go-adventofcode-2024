package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	l := ContentToLine(input)
	/*for c := 0; c < 25; c++ {
		fmt.Println(c)
		l = ApplyRules(l)
	}*/
	for n := 0; n <= 75; n += 5 {
		c := &sync.Map{}
		start := time.Now()
		fmt.Printf("n(%v): %v, time: %vs\n", n, RecursiveRules(l, n, 1, c), time.Since(start).Seconds())
	}
}
