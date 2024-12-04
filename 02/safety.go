package main

func safetyReport(reports []Report) int {
	r := 0
	for _, rep := range reports {
		if rep.IsSave() {
			r += 1
		}
	}
	return r
}
