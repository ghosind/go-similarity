package tokenizer_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/tokenizer"
)

func TestWhitespaceTokenizer(t *testing.T) {
	a := assert.New(t)
	tok := &tokenizer.WhitespaceTokenizer{}

	res, err := tok.Tokenize("A B  C")
	a.NilNow(err)
	a.Equal([]string{"A", "B", "C"}, res)
}
