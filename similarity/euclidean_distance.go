package similarity

import (
	"math"

	"github.com/ghosind/collection/set"
	"github.com/ghosind/go-similarity/tokenizer"
)

// EuclideanDistance algorithm calculates the Euclidean distance between two strings using the
// vector space of combined terms as the dimensions.
type EuclideanDistance struct {
	Tokenizer tokenizer.Tokenizer
}

// Compare returns the similarity score between two strings by Euclidean distance algorithm.
func (ed *EuclideanDistance) Compare(s1, s2 string) float64 {
	t1 := ed.getTokenizer().Tokenize(s1)
	t2 := ed.getTokenizer().Tokenize(s2)

	totalPossible := math.Sqrt(float64(len(t1))*float64(len(t1)) + float64(len(t2))*float64(len(t2)))
	if totalPossible == 0 {
		return 1
	}

	totalDistance := ed.euclideanDistance(t1, t2)

	return (totalPossible - totalDistance) / totalPossible
}

func (ed *EuclideanDistance) getTokenizer() tokenizer.Tokenizer {
	if ed.Tokenizer == nil {
		ed.Tokenizer = &tokenizer.WhitespaceTokenizer{}
	}
	return ed.Tokenizer
}

func (ed *EuclideanDistance) euclideanDistance(t1, t2 []string) float64 {
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	allTokens := set.NewHashSet[string]()
	allTokens.AddAll(t1...)
	allTokens.AddAll(t2...)

	for _, t := range t1 {
		m1[t]++
	}
	for _, t := range t2 {
		m2[t]++
	}

	totalDistance := 0.0
	for tok := range allTokens.Iter() {
		cnt1 := m1[tok]
		cnt2 := m2[tok]

		totalDistance += float64(cnt1-cnt2) * float64(cnt1-cnt2)
	}

	return math.Sqrt(totalDistance)
}
