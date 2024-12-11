package main

import "testing"

func Test_ContentToPuzzle(t *testing.T) {
	in := []byte(`AA
BB`)

	want := Puzzle{
		{
			Tile{C: 'A'}, Tile{C: 'A'},
		},
		{
			Tile{C: 'B'}, Tile{C: 'B'},
		},
	}
	got := ContentToPuzzle(in)
	for lid, l := range got {
		for eid, e := range l {
			if want[lid][eid].C != e.C {
				t.Errorf("wanted %#v, got %#v", want[lid][eid], e)
			}
		}
	}
}
