package main

import (
	"bytes"
	"math"
	"strconv"
)

type Report []float64

func inputToReports(content []byte) []Report {
	var r []Report
	lines := bytes.Split(content, []byte("\n"))

	for _, l := range lines {
		var rep Report
		values := bytes.Split(l, []byte(" "))
		for _, v := range values {
			f, _ := strconv.ParseFloat(string(v), 64)
			rep = append(rep, f)
		}
		r = append(r, rep)
	}

	return r
}

func (r Report) IsSave() bool {
	var up []float64
	var down []float64
	last := r[0]
	for _, v := range r[1:] {
		d := v - last
		if d > 0 {
			up = append(up, d)
		} else {
			down = append(down, d)
		}
		last = v
	}

	if len(up) > 0 && len(down) > 0 {
		return false // mixed steps
	}

	all := append(up, down...)
	for _, v := range all {
		a := math.Abs(v)
		if a < 1 || a > 3 {
			return false
		}
	}

	return true
}

func (r Report) IsSaveWithThreshold() bool {
	return true
}
