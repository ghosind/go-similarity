package similarity

import (
	"github.com/ghosind/go-similarity/cost_func"
)

// Levenshtein is the distance between two words is the minimum number of single-character edits
// (insertions, deletions or substitutions) required to change one word into the other.
type Levenshtein struct {
	CostFunc cost_func.SubstitutionCost
}

// Compare returns the Levenshtein similarity between two strings.
func (l *Levenshtein) Compare(s1, s2 string) (float64, error) {
	distance := l.getDistance(s1, s2)

	maxLen := float64(max(len(s1), len(s2)))

	if maxLen == 0 {
		return 1, nil
	} else {
		return 1 - distance/maxLen, nil
	}
}

// getDistance calculates the un-normalized Levenshtein distance between two strings.
func (l *Levenshtein) getDistance(s1, s2 string) float64 {
	n := len(s1)
	m := len(s2)
	if n == 0 {
		return float64(m)
	}
	if m == 0 {
		return float64(n)
	}

	matrix := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		matrix[i] = make([]float64, m+1)
		matrix[i][0] = float64(i)
	}
	for j := 0; j <= m; j++ {
		matrix[0][j] = float64(j)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			cost := l.getCostFunc().Cost(s1, i-1, s2, j-1)
			matrix[i][j] = min(
				matrix[i-1][j]+1,
				matrix[i][j-1]+1,
				matrix[i-1][j-1]+cost,
			)
		}
	}

	return matrix[n][m]
}

// getCostFunc returns the cost function to use for calculating the Levenshtein distance,
// defaulting to the SubCost01 cost function.
func (l *Levenshtein) getCostFunc() cost_func.SubstitutionCost {
	if l.CostFunc == nil {
		l.CostFunc = new(cost_func.SubCost01)
	}
	return l.CostFunc
}
