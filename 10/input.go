package main

import (
	"bytes"
	"strconv"
)

func ContentToMap(in []byte) Map {
	m := Map{}
	lines := bytes.Split(in, []byte("\n"))
	for y, l := range lines {
		var line []Level
		for x, v := range l {
			lvl, _ := strconv.ParseInt(string(v), 10, 64)
			line = append(line, Level{V: int(lvl), X: x, Y: y, M: &m})
		}
		m = append(m, line)
	}
	return m
}
