package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
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

func RecursiveRules(l Line, target, lvl int, cache *sync.Map) int {
	var ident strings.Builder
	ident.WriteString(fmt.Sprintf("%v_", lvl))
	for _, e := range l {
		ident.WriteString(fmt.Sprintf(`%v;`, e.ToString()))
	}

	if cache != nil {
		if v, ok := cache.Load(ident.String()); ok {
			return v.(int)
		}
	}

	if target == lvl-1 {
		return len(l)
	}

	n := 0
	for _, e := range l {
		switch {
		case e.Numeric == 0:
			n += RecursiveRules(Line{
				&Element{
					Numeric: 1,
				},
			}, target, lvl+1, nil)
		case len(e.ToString())%2 == 0:
			s := e.ToString()
			x := len(s) / 2
			first, _ := strconv.ParseInt(s[0:x], 10, 64)
			second, _ := strconv.ParseInt(s[x:], 10, 64)
			n += RecursiveRules(Line{
				&Element{
					Numeric: int(first),
				}, &Element{
					Numeric: int(second),
				},
			}, target, lvl+1, cache)
		default:
			n += RecursiveRules(Line{
				&Element{
					Numeric: e.Numeric * 2024,
				},
			}, target, lvl+1, cache)
		}
	}

	if cache != nil {
		cache.Store(ident.String(), n)
	}
	return n
}
