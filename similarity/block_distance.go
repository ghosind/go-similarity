package similarity

import (
	"github.com/ghosind/collection/set"
	"github.com/ghosind/go-similarity/tokenizer"
)

// BlockDistance algorithm is a similarity algorithm that calculates the similarity between two
// strings by comparing the number of tokens that are not in common between the two strings.
type BlockDistance struct {
	Tokenizer tokenizer.Tokenizer
}

// Compare returns the similarity score between two strings by block distance algorithm.
func (d *BlockDistance) Compare(s1, s2 string) (float64, error) {
	t1, err := d.getTokenizer().Tokenize(s1)
	if err != nil {
		return 0, err
	}
	t2, err := d.getTokenizer().Tokenize(s2)
	if err != nil {
		return 0, err
	}

	totalPossible := float64(len(t1) + len(t2))
	if totalPossible == 0 {
		return 1, nil
	}

	totalDistance := d.getUnNormalizedSimilarity(t1, t2)

	return (totalPossible - totalDistance) / totalPossible, nil
}

func (d *BlockDistance) getTokenizer() tokenizer.Tokenizer {
	if d.Tokenizer == nil {
		d.Tokenizer = &tokenizer.WhitespaceTokenizer{}
	}
	return d.Tokenizer
}

func (d *BlockDistance) getUnNormalizedSimilarity(t1, t2 []string) float64 {
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	for _, s := range t1 {
		m1[s]++
	}
	for _, s := range t2 {
		m2[s]++
	}

	s := set.NewHashSet[string]()
	s.AddAll(t1...)
	s.AddAll(t2...)

	totalDistance := 0
	for tok := range s.Iter() {
		cntS1 := m1[tok]
		cntS2 := m2[tok]

		if cntS1 > cntS2 {
			totalDistance += cntS1 - cntS2
		} else {
			totalDistance += cntS2 - cntS1
		}
	}

	return float64(totalDistance)
}
