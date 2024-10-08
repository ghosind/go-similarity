package tokenizer_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/tokenizer"
)

func TestWhitespaceTokenizer(t *testing.T) {
	a := assert.New(t)
	tok := &tokenizer.WhitespaceTokenizer{}

	res := tok.Tokenize("A B  C")
	a.Equal([]string{"A", "B", "C"}, res)
}
