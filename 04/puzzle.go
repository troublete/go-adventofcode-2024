package main

import (
	"strings"
)

type Puzzle [][]Tile

func (p *Puzzle) Get(x, y int) *Tile {
	if y >= 0 && y <= len(*p)-1 {
		if x >= 0 && x <= len((*p)[y])-1 {
			return &(*p)[y][x]
		}

	}
	return nil
}

func (p *Puzzle) FindAll(c byte) []*Tile {
	var f []*Tile
	for _, l := range *p {
		for _, e := range l {
			if e.C == c {
				f = append(f, &e)
			}
		}
	}
	return f
}

type Tile struct {
	C      byte
	X, Y   int
	Puzzle *Puzzle
}

type Change struct {
	X, Y int
}

func (t *Tile) CheckFor(term string) (int, []*Tile) {
	chars := strings.Split(term, "")
	directions := []Change{
		{1, 0},   // to right
		{-1, 0},  // to left
		{0, 1},   // to bottom
		{0, -1},  // to top
		{1, 1},   // to bottom right
		{1, -1},  // to top right
		{-1, -1}, // to top left
		{-1, 1},  // to bottom left
	}

	n := 0
	var affected []*Tile

	for _, d := range directions {
		found := chars[0]
		tiles := []*Tile{t}
		x, y := t.X, t.Y
		for range chars[1:] {
			x += d.X
			y += d.Y

			adjacent := t.Puzzle.Get(x, y)
			if adjacent != nil {
				found += string(adjacent.C)
				tiles = append(tiles, adjacent)
			}
		}

		if found == term {
			n += 1
			affected = append(affected, tiles...)
		}
	}
	return n, affected
}
