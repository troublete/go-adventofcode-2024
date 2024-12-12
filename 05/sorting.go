package main

type Sorting map[int][]int

type Manual []int

func (m Manual) MiddlePage() int {
	rest := len(m) % 2
	return m[((len(m) - rest) / 2)]
}

type Manuals []Manual

type Validator struct {
	S Sorting
	M Manuals
}

func (v *Validator) ValidManuals() []Manual {
	var valid []Manual
	for _, manual := range v.M {
		manualValid := true
		for did, d := range manual {
			before := manual[:did]
			if requirement, ok := v.S[d]; ok { // if rule for number exists
				for _, r := range requirement { // check 'after' requirement
					for _, b := range before { // on the before's
						if b == r { // if there is a match; mark manual as invalid
							manualValid = false
						}
					}
				}
			}
		}
		if manualValid {
			valid = append(valid, manual)
		}
	}
	return valid
}
