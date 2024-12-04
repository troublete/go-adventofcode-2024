package main

import "testing"

func Test_similarity(t *testing.T) {
	a := []float64{3, 4, 2, 1, 3, 3}
	b := []float64{4, 3, 5, 3, 9, 3}
	if similarity(a, b) != 31 {
		t.Errorf("expected 31, got %v", similarity(a, b))
	}
}
