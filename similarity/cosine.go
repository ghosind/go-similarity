package similarity

import (
	"math"

	"github.com/ghosind/collection/set"
	"github.com/ghosind/go-similarity/tokenizer"
)

// Cosine is a similarity measure that calculates the cosine of the angle between two vectors.
type Cosine struct {
	Tokenizer tokenizer.Tokenizer
}

// Compare returns the similarity score between two strings by cosine similarity algorithm.
func (c *Cosine) Compare(s1, s2 string) float64 {
	t1 := c.getTokenizer().Tokenize(s1)
	t2 := c.getTokenizer().Tokenize(s2)
	if len(t1) == 0 && len(t2) == 0 {
		return 1
	}

	allTokens := set.NewHashSet[string]()
	allTokens.AddAll(t1...)
	terms1 := allTokens.Size()
	s2Tokens := set.NewHashSet[string]()
	s2Tokens.AddAll(t2...)
	terms2 := s2Tokens.Size()

	allTokens.AddAll(t2...)
	commonTerms := (terms1 + terms2) - allTokens.Size()

	return float64(commonTerms) / (math.Sqrt(float64(terms1)) * math.Sqrt(float64(terms2)))
}

func (c *Cosine) getTokenizer() tokenizer.Tokenizer {
	if c.Tokenizer == nil {
		c.Tokenizer = &tokenizer.WhitespaceTokenizer{}
	}
	return c.Tokenizer
}
