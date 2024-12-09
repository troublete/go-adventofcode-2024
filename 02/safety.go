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

func safetyReportWithTolerance(reports []Report) int {
	r := 0
	for _, rep := range reports {
		if rep.IsSaveWithThreshold() {
			r += 1
		}
	}
	return r
}
