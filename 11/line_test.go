package main

import (
	"testing"
)

func Test_ApplyRules(t *testing.T) {
	e := ContentToLine(ex)
	e = ApplyRules(e)
	want := []int{253000, 1, 7}
	for wid, w := range want {
		if e[wid].Numeric != w {
			t.Error("failed to generate the correct numbers")
		}
	}

	e = ApplyRules(e)
	want = []int{253, 0, 2024, 14168}
	for wid, w := range want {
		if e[wid].Numeric != w {
			t.Error("failed to generate the correct numbers")
		}
	}

}
