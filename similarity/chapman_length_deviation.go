package similarity

// ChapmanLengthDeviation algorithm calculates the length deviation of two strings.
type ChapmanLengthDeviation struct{}

// Compare returns the similarity score between two strings by Chapman length deviation algorithm.
func (s *ChapmanLengthDeviation) Compare(s1, s2 string) (float64, error) {
	if len(s1) == 0 && len(s2) == 0 {
		return 1, nil
	}

	if len(s1) > len(s2) {
		return float64(len(s2)) / float64(len(s1)), nil
	} else {
		return float64(len(s1)) / float64(len(s2)), nil
	}
}
