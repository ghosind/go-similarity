package cost_func

type SubCost1Minus2 struct{}

func (f *SubCost1Minus2) Cost(s1 string, i1 int, s2 string, i2 int) float64 {
	if len(s1) <= i1 || len(s2) <= i2 {
		return 0
	}

	if s1[i1] == s2[i2] {
		return f.GetMaxCost()
	} else {
		return f.GetMinCost()
	}
}

func (f *SubCost1Minus2) GetMaxCost() float64 {
	return 1
}

func (f *SubCost1Minus2) GetMinCost() float64 {
	return -2
}
