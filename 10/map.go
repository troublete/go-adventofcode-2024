package main

import "fmt"

type Level struct {
	V, X, Y int
	M       *Map
}

func (l *Level) FindPathTo(v int) []*Level {
	if l.V == v {
		return []*Level{l}
	}

	toCheck := []*Level{
		l.M.Get(l.X, l.Y-1), // top
		l.M.Get(l.X+1, l.Y), // right
		l.M.Get(l.X-1, l.Y), // left
		l.M.Get(l.X, l.Y+1), // bottom
	}

	var ends []*Level
	for _, t := range toCheck {
		if t != nil && t.V == l.V-1 {
			if l.V != v {
				te := t.FindPathTo(v)
				ends = append(ends, te...)
			}
		}
	}
	return ends
}

type Path struct {
	Start, End *Level
}

func (m Map) FindAllPaths(start, end int) []Path {
	u := map[string]Path{}
	for _, s := range m.FindAll(end) {
		ends := s.FindPathTo(start)
		for _, e := range ends {
			ident := fmt.Sprintf("%v;%v|%v;%v", s.X, s.Y, e.X, e.Y)
			if _, ok := u[ident]; !ok {
				u[ident] = Path{
					Start: s,
					End:   e,
				}
			}
		}
	}

	var p []Path
	for _, path := range u {
		p = append(p, path)
	}

	return p
}

type Map [][]Level

func (m Map) FindAll(x int) []*Level {
	var ends []*Level
	for _, l := range m {
		for _, v := range l {
			if v.V == x {
				ends = append(ends, &v)
			}
		}
	}
	return ends
}

func (m Map) Get(x, y int) *Level {
	if y >= 0 && y <= len(m)-1 {
		if x >= 0 && x <= len(m[y])-1 {
			return &m[y][x]
		}
	}
	return nil
}
