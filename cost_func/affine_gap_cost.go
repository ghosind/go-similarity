package cost_func

type AffineGapCost interface {
	Cost(s string, startGap, endGap int) float64
}