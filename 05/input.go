package main

import (
	"bytes"
	"regexp"
	"strconv"
)

var (
	ReOrder  = regexp.MustCompile(`^(\d+)\|(\d+)$`)
	ReManual = regexp.MustCompile(`^(\d+,?)+$`)
)

func ContentToValidator(in []byte) *Validator {
	v := Validator{
		S: map[int][]int{},
		M: []Manual{},
	}
	for _, l := range bytes.Split(in, []byte("\n")) {
		switch {
		case ReOrder.Match(l):
			m := ReOrder.FindSubmatch(l)
			before, _ := strconv.ParseInt(string(m[1]), 10, 64)
			after, _ := strconv.ParseInt(string(m[2]), 10, 64)
			if cond, ok := v.S[int(before)]; ok {
				v.S[int(before)] = append(cond, int(after))
			} else {
				v.S[int(before)] = []int{int(after)}
			}
		case ReManual.Match(l):
			m := regexp.MustCompile(`(\d+)`).FindAllSubmatch(l, -1)
			var ml []int
			for _, sm := range m {
				x, _ := strconv.ParseInt(string(sm[1]), 10, 64)
				ml = append(ml, int(x))
			}
			v.M = append(v.M, ml)
		}
	}

	return &v
}
