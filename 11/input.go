package main

import (
	"bytes"
	"strconv"
)

func ContentToLine(in []byte) Line {
	var l Line

	parts := bytes.Split(in, []byte(" "))
	for _, p := range parts {
		n, _ := strconv.ParseInt(string(p), 10, 64)
		l = append(l, &Element{
			Numeric:    int(n),
			HasChanged: false,
		})
	}

	return l
}
