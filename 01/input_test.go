package main

import (
	"testing"
)
import _ "embed"

//go:embed example.txt
var example []byte

func Test_inputToLists(t *testing.T) {
	a, b := inputToLists(example)

	aCheck := []float64{3, 4, 2, 1, 3, 3}
	bCheck := []float64{4, 3, 5, 3, 9, 3}

	check := func(want, got []float64) {
		for idx, w := range want {
			if got[idx] != w {
				t.Errorf("expected %v, got %v", w, got[idx])
			}
		}
	}

	check(aCheck, a)
	check(bCheck, b)
}
