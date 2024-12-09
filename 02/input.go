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

func (r Report) Clean() Report {
	last := r[0]
	cleaned := Report{last}
	var asc, desc []float64

	removeIdx := func(in []float64, idx int) []float64 {
		var after []float64
		if len(in)-1 < idx {
			after = []float64{}
		} else {
			after = in[idx+1:]
		}
		return append(in[:idx], after...)
	}

	for idx, v := range r[1:] {
		d := v - last
		absD := math.Abs(d)
		if absD < 1 || absD > 3 {
			cleaned = removeIdx(cleaned, idx)
		}

		if d >= 0 {
			if len(desc) > 0 {
				cleaned = removeIdx(cleaned, idx)
			} else {
				asc = append(asc, v)
			}
		} else {
			if len(asc) > 0 {
				cleaned = removeIdx(cleaned, idx)
			} else {
				desc = append(desc, v)
			}
		}

		cleaned = append(cleaned, v)
		last = v
	}

	return cleaned
}

func (r Report) save() bool {
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

func (r Report) IsSave() bool {
	return r.save()
}

func (r Report) IsSaveWithThreshold() bool {
	return r.Clean().save()
}
