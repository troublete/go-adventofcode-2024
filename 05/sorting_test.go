package main

import (
	"testing"
)

func TestValidManual(t *testing.T) {
	v := ContentToValidator(ex1)
	vm := v.ValidManuals()

	want := []int{61, 53, 29}
	for mid, m := range vm {
		got := m.MiddlePage()
		if want[mid] != got {
			t.Errorf("wanted %v, got %v", want[mid], got)
		}
	}
}
