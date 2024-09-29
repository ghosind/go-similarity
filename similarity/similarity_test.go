package similarity_test

import "math"

const (
	epsilon = 1e-9
)

func isSameFloat64(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
