package cost_function

type AffineGapCost interface {
	Cost(s string, startGap, endGap int) float64
}
