package main

import "testing"
import _ "embed"

//go:embed example.txt
var example []byte

func Test_inputToReports(t *testing.T) {
	r := inputToReports(example)
	want := []Report{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	for wid, w := range want {
		for vid, v := range w {
			if r[wid][vid] != v {
				t.Errorf("wanted %v, got %v", v, r[wid][vid])
			}
		}
	}
}

func Test_Save(t *testing.T) {
	r := inputToReports(example)
	want := []bool{
		true,
		false,
		false,
		false,
		false,
		true,
	}

	for rid, rep := range r {
		if want[rid] != rep.IsSave() {
			t.Errorf("wanted %v, got %v", want[rid], rep.IsSave())
		}
	}
}

func Test_SaveWithThreshold(t *testing.T) {
	r := inputToReports(example)
	want := []bool{
		true,
		false,
		false,
		true,
		true,
		true,
	}

	for rid, rep := range r {
		if want[rid] != rep.IsSaveWithThreshold(1) {
			t.Errorf("#%v: wanted %v, got %v", rid, want[rid], rep.IsSaveWithThreshold(1))
		}
	}
}
