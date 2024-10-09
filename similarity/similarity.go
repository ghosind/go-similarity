package similarity

// Similarity is an interface for comparing similarity of two strings.
type Similarity interface {
	// Compare returns the similarity score between two strings.
	Compare(s1, s2 string) (float64, error)
}
