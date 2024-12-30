package similarity

import "github.com/ghosind/go-similarity/cost_func"

// NeedlemanWunsch is a similarity algorithm that calculates the similarity between two strings
// using the Needleman-Wunsch algorithm.
type NeedlemanWunsch struct {
	CostFunc cost_func.SubstitutionCost
	GapCost  float64
}

// Compare returns the Needleman-Wunsch similarity between two strings.
func (n *NeedlemanWunsch) Compare(s1, s2 string) (float64, error) {
	distance := n.getDistance(s1, s2)

	maxValue := float64(max(len(s1), len(s2)))
	minValue := maxValue

	if n.getCostFunc().GetMaxCost() > n.getGapCost() {
		maxValue *= n.getCostFunc().GetMaxCost()
	} else {
		maxValue *= n.getGapCost()
	}
	if n.getCostFunc().GetMinCost() < n.getGapCost() {
		minValue *= n.getCostFunc().GetMinCost()
	} else {
		minValue *= n.getGapCost()
	}

	if minValue < 0 {
		maxValue -= minValue
		distance -= minValue
	}

	if maxValue == 0 {
		return 1.0, nil
	} else {
		return 1.0 - distance/maxValue, nil
	}
}

// getDistance calculates the distance between two strings using the Needleman-Wunsch algorithm.
func (n *NeedlemanWunsch) getDistance(s1, s2 string) float64 {
	l1 := len(s1)
	l2 := len(s2)

	if l1 == 0 {
		return float64(l2)
	}
	if l2 == 0 {
		return float64(l1)
	}

	matrix := make([][]float64, l1+1)
	for i := range matrix {
		matrix[i] = make([]float64, l2+1)
		matrix[i][0] = float64(i)
	}
	for j := range matrix[0] {
		matrix[0][j] = float64(j)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			cost := n.getCostFunc().Cost(s1, i-1, s2, j-1)

			matrix[i][j] = min(
				matrix[i-1][j]+n.getGapCost(),
				matrix[i][j-1]+n.getGapCost(),
				matrix[i-1][j-1]+cost,
			)
		}
	}

	return matrix[l1][l2]
}

// getCostFunc returns the cost function to use for calculating the Needleman-Wunsch algorithm,
// defaulting to the SubCost01 cost function.
func (n *NeedlemanWunsch) getCostFunc() cost_func.SubstitutionCost {
	if n.CostFunc == nil {
		n.CostFunc = &cost_func.SubCost01{}
	}
	return n.CostFunc
}

// getGapCost returns the cost of a gap, defaulting to 2.0.
func (n *NeedlemanWunsch) getGapCost() float64 {
	if n.GapCost == 0 {
		n.GapCost = 2.0
	}
	return n.GapCost
}
