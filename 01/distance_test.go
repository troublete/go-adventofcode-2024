package main

import "testing"

func Test_distance(t *testing.T) {
	a := []float64{3, 4, 2, 1, 3, 3}
	b := []float64{4, 3, 5, 3, 9, 3}
	if distance(a, b) != 11 {
		t.Errorf("expected 11, got %v", distance(a, b))
	}
}
