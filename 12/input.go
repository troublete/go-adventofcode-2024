package main

import "bytes"

func ContentToMap(in []byte) *Map {
	var m Map
	lines := bytes.Split(in, []byte("\n"))
	for y, l := range lines {
		var li []*Tile
		for x, e := range l {
			li = append(li, &Tile{
				C:   e,
				X:   x,
				Y:   y,
				Map: &m,
			})
		}
		m = append(m, li)
	}
	return &m
}
