package main

import (
	_ "embed"
	"math"
	"sort"
)

func distance(a, b []float64) float64 {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	dist := 0.0
	for idx, v := range a {
		dist += math.Abs(v - b[idx])
	}

	return dist
}
