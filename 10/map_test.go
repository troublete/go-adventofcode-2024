package main

import (
	"testing"
)

func Test_AllEnds(t *testing.T) {
	m := ContentToMap(ex)
	if len(m.FindAll(9)) != 1 {
		t.Error("should be just one end")
	}
}

func Test_FindAllPaths(t *testing.T) {
	m := ContentToMap(ex2)
	p := m.FindAllPaths(0, 9)
	if len(p) != 36 {
		t.Error("wrong amount of paths found")
	}
}
