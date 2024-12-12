package main

import "testing"

func TestContentToMap(t *testing.T) {
	got := ContentToMap(ex1)
	want := Map{
		{&Tile{C: 'A'}, &Tile{C: 'A'}, &Tile{C: 'A'}, &Tile{C: 'A'}},
		{&Tile{C: 'B'}, &Tile{C: 'B'}, &Tile{C: 'C'}, &Tile{C: 'D'}},
		{&Tile{C: 'B'}, &Tile{C: 'B'}, &Tile{C: 'C'}, &Tile{C: 'C'}},
		{&Tile{C: 'E'}, &Tile{C: 'E'}, &Tile{C: 'E'}, &Tile{C: 'C'}},
	}

	for lid, l := range want {
		for eid, e := range l {
			if e.C != got.Get(eid, lid).C {
				t.Errorf("wanted %#v, got %#v", e, got.Get(eid, lid))
			}
		}
	}
}
