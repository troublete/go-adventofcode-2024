package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_FindAllOccurences(t *testing.T) {
	p := ContentToPuzzle(example)
	xs := p.FindAll('X')
	n := 0
	var highlighted []*Tile
	for _, o := range xs {
		x, tiles := o.CheckFor("XMAS")
		n += x
		highlighted = append(highlighted, tiles...)
	}

	highlightMap := map[string]*Tile{}
	for _, h := range highlighted {
		ident := fmt.Sprintf("%v_%v", h.X, h.Y)
		highlightMap[ident] = h
	}

	out := bytes.Buffer{}
	for _, l := range p {
		for _, e := range l {
			ident := fmt.Sprintf("%v_%v", e.X, e.Y)
			if _, ok := highlightMap[ident]; ok {
				out.Write([]byte{e.C})
			} else {
				out.WriteString(".")
			}
		}
		out.WriteString("\n")
	}

	if n != 18 {
		t.Log(out.String())
		t.Error("wrong amount of words found")
	}
}
