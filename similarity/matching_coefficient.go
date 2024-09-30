package similarity

import (
	"github.com/ghosind/collection/set"
	"github.com/ghosind/go-similarity/tokenizer"
)

// MatchingCoefficient is a similarity algorithm that measures the similarity between two strings.
type MatchingCoefficient struct {
	Tokenizer tokenizer.Tokenizer
}

// Compare returns the Matching Coefficient similarity between two strings.
func (mc *MatchingCoefficient) Compare(s1, s2 string) float64 {
	t1 := mc.getTokenizer().Tokenize(s1)
	t2 := mc.getTokenizer().Tokenize(s2)
	if len(t1) == 0 && len(t2) == 0 {
		return 1
	}

	totalPossible := max(len(t1), len(t2))
	totalFound := mc.getUnNormalizedSimilarity(t1, t2)

	return float64(totalFound) / float64(totalPossible)
}

func (mc *MatchingCoefficient) getUnNormalizedSimilarity(t1, t2 []string) int {
	m1 := set.NewHashSet[string]()
	m1.AddAll(t1...)
	m2 := set.NewHashSet[string]()
	m2.AddAll(t2...)

	totalFound := 0
	for tok := range m1.Iter() {
		if m2.Contains(tok) {
			totalFound++
		}
	}

	return totalFound
}

func (mc *MatchingCoefficient) getTokenizer() tokenizer.Tokenizer {
	if mc.Tokenizer == nil {
		mc.Tokenizer = &tokenizer.WhitespaceTokenizer{}
	}
	return mc.Tokenizer
}
