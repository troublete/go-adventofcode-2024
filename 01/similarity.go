package main

func similarity(a, b []float64) float64 {
	bLookup := map[float64]int{}
	for _, bv := range b {
		if v, ok := bLookup[bv]; !ok {
			bLookup[bv] = 1
		} else {
			bLookup[bv] = v + 1
		}
	}

	s := 0.0
	for _, av := range a {
		if count, ok := bLookup[av]; ok {
			s += av * float64(count)
		}
	}

	return s
}
