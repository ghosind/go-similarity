package cost_function

import "github.com/ghosind/collection/set"

func init() {
	subCost5_3Minus3_approx = make([]*set.HashSet[byte], len(subCost5_3Minus3_approx_chars))

	for i, chars := range subCost5_3Minus3_approx_chars {
		subCost5_3Minus3_approx[i] = set.NewHashSet[byte]()
		subCost5_3Minus3_approx[i].AddAll(chars...)
	}
}
