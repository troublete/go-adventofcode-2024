package main

import "bytes"

func ContentToPuzzle(in []byte) Puzzle {
	var p Puzzle
	lines := bytes.Split(in, []byte("\n"))
	for y, l := range lines {
		var li []Tile
		for x, e := range l {
			li = append(li, Tile{
				C:      e,
				X:      x,
				Y:      y,
				Puzzle: &p,
			})
		}
		p = append(p, li)
	}
	return p
}
