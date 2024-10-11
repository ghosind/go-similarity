package cost_function

type SubCost01 struct{}

func (f *SubCost01) Cost(s1 string, i1 int, s2 string, i2 int) float64 {
	if len(s1) <= i1 || len(s2) <= i2 {
		return 1
	}

	if s1[i1] == s2[i2] {
		return 0
	} else {
		return 1
	}
}
