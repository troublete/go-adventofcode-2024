package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	Numeric    int
	HasChanged bool
}

func (e Element) ToString() string {
	return fmt.Sprintf("%d", e.Numeric)
}

type Line []*Element

func ApplyRules(l Line) Line {
	toProcess := len(l) - 1
	for eid := 0; eid <= toProcess; eid++ {
		e := l[eid]
		if e.HasChanged {
			continue
		}

		switch {
		case e.Numeric == 0:
			e.Numeric = 1
			e.HasChanged = true
		case len(e.ToString())%2 == 0:
			s := e.ToString()
			n := len(s) / 2
			first, _ := strconv.ParseInt(s[0:n], 10, 64)
			second, _ := strconv.ParseInt(s[n:], 10, 64)
			after := Line{}
			for _, ele := range l[eid+1:] {
				after = append(after, &Element{
					Numeric:    ele.Numeric,
					HasChanged: false,
				})
			}
			l = append(l[:eid], &Element{
				Numeric:    int(first),
				HasChanged: true,
			}, &Element{
				Numeric:    int(second),
				HasChanged: true,
			})
			l = append(l, after...)
			toProcess++
		default:
			e.Numeric = e.Numeric * 2024
			e.HasChanged = true
		}
	}

	l.Reset()
	return l
}

func (l Line) Reset() {
	for _, e := range l {
		e.HasChanged = false
	}
}
