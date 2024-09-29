package tokenizer

import "strings"

// WhitespaceTokenizer is a tokenizer that splits the input string by whitespace.
type WhitespaceTokenizer struct {
	Tokenizer
}

func (t *WhitespaceTokenizer) Tokenize(input string) []string {
	return strings.Fields(input)
}
