package similarity

import (
	"github.com/ghosind/collection/set"
	"github.com/ghosind/go-similarity/tokenizer"
)

// Jaccard is a similarity algorithm that measures the similarity between two strings.
type Jaccard struct {
	Tokenizer tokenizer.Tokenizer
}

// Compare returns the Jaccard similarity between two strings.
func (j *Jaccard) Compare(s1, s2 string) float64 {
	t1 := j.getTokenizer().Tokenize(s1)
	t2 := j.getTokenizer().Tokenize(s2)
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

	return float64(commonTerms) / float64(allTokens.Size())
}

func (j *Jaccard) getTokenizer() tokenizer.Tokenizer {
	if j.Tokenizer == nil {
		j.Tokenizer = &tokenizer.WhitespaceTokenizer{}
	}
	return j.Tokenizer
}
