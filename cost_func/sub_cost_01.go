package cost_func

type SubCost01 struct{}

func (f *SubCost01) Cost(s1 string, i1 int, s2 string, i2 int) float64 {
	if len(s1) <= i1 || len(s2) <= i2 {
		return f.GetMaxCost()
	}

	if s1[i1] == s2[i2] {
		return f.GetMinCost()
	} else {
		return f.GetMaxCost()
	}
}

func (f *SubCost01) GetMaxCost() float64 {
	return 1.0
}

func (f *SubCost01) GetMinCost() float64 {
	return 0.0
}
