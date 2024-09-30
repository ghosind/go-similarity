// go:build !go1.21

package similarity

import "cmp"

func max[T cmp.Ordered](n T, nums ...T) T {
	if len(nums) == 0 {
		return n
	}

	maxN := n
	for _, num := range nums {
		if num > maxN {
			maxN = num
		}
	}
	return maxN
}

func min[T cmp.Ordered](n T, nums ...T) T {
	if len(nums) == 0 {
		return n
	}

	minN := n
	for _, num := range nums {
		if num < minN {
			minN = num
		}
	}
	return minN
}
