package main

import "testing"
import _ "embed"

//go:embed example.txt
var example []byte

//go:embed example_2.txt
var secondExample []byte

func Test_Extract(t *testing.T) {
	calcs := Extract(example)
	for idx, c := range []Calc{
		{
			A: 2,
			B: 4,
		},
		{
			A: 5,
			B: 5,
		},
		{
			A: 11,
			B: 8,
		},
		{
			A: 8,
			B: 5,
		},
	} {
		if idx > len(calcs)-1 {
			t.Error("to few extractions")
		}

		if c.A != calcs[idx].A || c.B != calcs[idx].B {
			t.Errorf("wrong arguments; want %#v, got %#v", c, calcs[idx])
		}
	}
}

func Test_SectionedExtract(t *testing.T) {
	sum := 0.0
	for _, e := range Extract(Interpret(secondExample)) {
		sum += e.Mul()
	}
	if sum != 48 {
		t.Error("wrong sum")
	}
}
