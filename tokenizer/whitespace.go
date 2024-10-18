package tokenizer

import "strings"

// WhitespaceTokenizer is a tokenizer that splits the input string by whitespace.
type WhitespaceTokenizer struct{}

func (t *WhitespaceTokenizer) Tokenize(input string) ([]string, error) {
	return strings.Fields(input), nil
}
