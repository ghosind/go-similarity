package cost_func

type SubstitutionCost interface {
	Cost(s1 string, i1 int, s2 string, i2 int) float64
}
