package main

import (
	"regexp"
	"strconv"
)

type Calc struct {
	A, B float64
}

func (c Calc) Mul() float64 {
	return c.A * c.B
}

func Extract(in []byte) []Calc {
	re := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))`)
	m := re.FindAllSubmatch(in, -1)
	var cs []Calc
	for _, sm := range m {
		a, _ := strconv.ParseFloat(string(sm[2]), 64)
		b, _ := strconv.ParseFloat(string(sm[3]), 64)
		cs = append(cs, Calc{
			A: a,
			B: b,
		})
	}
	return cs
}

func Interpret(in []byte) []byte {
	var contents []byte
	var op []byte
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)
	enabled := true
	for _, b := range in {
		op = append(op, b)
		switch {
		case reDo.Match(op):
			enabled = true
			op = []byte{}
		case reDont.Match(op):
			enabled = false
			op = []byte{}
		}

		if enabled {
			contents = append(contents, b)
		}
	}

	return contents
}
