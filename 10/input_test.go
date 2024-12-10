package main

import "testing"
import _ "embed"

//go:embed example.txt
var ex []byte

//go:embed example2.txt
var ex2 []byte

func Test_ContentToMap(t *testing.T) {
	m := ContentToMap(ex)
	want := Map{
		{Level{V: 0}, Level{V: 1}, Level{V: 2}, Level{V: 3}},
		{Level{V: 1}, Level{V: 2}, Level{V: 3}, Level{V: 4}},
		{Level{V: 8}, Level{V: 7}, Level{V: 6}, Level{V: 5}},
		{Level{V: 9}, Level{V: 8}, Level{V: 7}, Level{V: 6}},
	}

	for lid, l := range m {
		for vid, v := range l {
			if want[lid][vid].V != v.V {
				t.Errorf("want %v, got %v (%v, %v)", want[lid][vid].V, v.V, lid, vid)
			}
		}
	}
}
