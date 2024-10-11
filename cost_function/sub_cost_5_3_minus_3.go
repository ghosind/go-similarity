package cost_function

import "github.com/ghosind/collection/set"

const (
	subCost5_3Minus3_Exact_Match  = 5.0
	subCost5_3Minus3_Approx_Match = 3.0
	subCost5_3Minus3_Mismatch     = -3.0
)

var subCost5_3Minus3_approx []*set.HashSet[byte]
var subCost5_3Minus3_approx_chars [][]byte = [][]byte{
	{'d', 't'},
	{'g', 'j'},
	{'l', 'r'},
	{'m', 'n'},
	{'b', 'p', 'v'},
	{'a', 'e', 'i', 'o', 'u'},
	{',', '.'},
}

type SubCost5_3Minus3 struct{}

func (f *SubCost5_3Minus3) Cost(s1 string, i1 int, s2 string, i2 int) float64 {
	if len(s1) <= i1 || len(s2) <= i2 {
		return subCost5_3Minus3_Mismatch
	}

	if s1[i1] == s2[i2] {
		return subCost5_3Minus3_Exact_Match
	}
	for _, approx := range subCost5_3Minus3_approx {
		if approx.Contains(s1[i1]) && approx.Contains(s2[i2]) {
			return subCost5_3Minus3_Approx_Match
		}
	}
	return subCost5_3Minus3_Mismatch
}
