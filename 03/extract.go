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

func Sections(in []byte) [][]byte {
	re := regexp.MustCompile(`(do\(\))?.*don\'t\(\)|do\(\).*(don\'t\(\))?`)
	m := re.FindAllSubmatch(in, -1)

	var sections [][]byte
	for _, sm := range m {
		sections = append(sections, sm[0])
	}
	return sections
}
