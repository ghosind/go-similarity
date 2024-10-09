package similarity

import (
	"github.com/ghosind/collection/set"
	"github.com/ghosind/go-similarity/tokenizer"
)

// OverlapCoefficient is a similarity algorithm that measures the similarity between two strings
// where the similarity is the size of the intersection divided by the size of the smallest set.
type OverlapCoefficient struct {
	Tokenizer tokenizer.Tokenizer
}

// Compare returns the Overlap Coefficient similarity between two strings.
func (o *OverlapCoefficient) Compare(s1, s2 string) (float64, error) {
	t1, err := o.getTokenizer().Tokenize(s1)
	if err != nil {
		return 0, err
	}
	t2, err := o.getTokenizer().Tokenize(s2)
	if err != nil {
		return 0, err
	}

	if len(t1) == 0 && len(t2) == 0 {
		return 1, nil
	}

	allTokens := set.NewHashSet[string]()
	allTokens.AddAll(t1...)
	terms1 := allTokens.Size()
	s2Tokens := set.NewHashSet[string]()
	s2Tokens.AddAll(t2...)
	terms2 := s2Tokens.Size()

	allTokens.AddAll(t2...)
	commonTerms := (terms1 + terms2) - allTokens.Size()

	return float64(commonTerms) / float64(min(terms1, terms2)), nil
}

func (o *OverlapCoefficient) getTokenizer() tokenizer.Tokenizer {
	if o.Tokenizer == nil {
		o.Tokenizer = &tokenizer.WhitespaceTokenizer{}
	}
	return o.Tokenizer
}
