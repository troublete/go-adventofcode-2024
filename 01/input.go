package main

import (
	"bytes"
	"strconv"
)

func inputToLists(content []byte) ([]float64, []float64) {
	var a, b []float64
	lines := bytes.Split(content, []byte("\n"))
	for _, l := range lines {
		v := bytes.Split(l, []byte("   "))
		av, _ := strconv.ParseFloat(string(v[0]), 64)
		bv, _ := strconv.ParseFloat(string(v[1]), 64)
		a = append(a, av)
		b = append(b, bv)
	}
	return a, b
}
