package similarity

import (
	"github.com/ghosind/collection/set"
	"github.com/ghosind/go-similarity/tokenizer"
)

// Dice algorithm provides a similarity measure between two strings using the vector space of
// present terms.
type Dice struct {
	Tokenizer tokenizer.Tokenizer
}

// Compare returns the similarity score between two strings by dice similarity algorithm.
func (d *Dice) Compare(s1, s2 string) float64 {
	t1 := d.getTokenizer().Tokenize(s1)
	t2 := d.getTokenizer().Tokenize(s2)
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

	return float64(commonTerms*2.0) / float64(terms1+terms2)
}

func (d *Dice) getTokenizer() tokenizer.Tokenizer {
	if d.Tokenizer == nil {
		d.Tokenizer = &tokenizer.WhitespaceTokenizer{}
	}
	return d.Tokenizer
}
