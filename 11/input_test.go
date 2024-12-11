package main

import "testing"

func Test_ContentToLine(t *testing.T) {
	l := ContentToLine(ex)
	if len(l) != 2 {
		t.Error("wanted to elements")
	}

	if l[0].Numeric != 125 || l[1].Numeric != 17 {
		t.Error("numbers parsed wrong")
	}
}
