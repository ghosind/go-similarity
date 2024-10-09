package tokenizer

// Tokenizer is an interface for tokenizing a string.
type Tokenizer interface {
	// Tokenize splits the input string into tokens.
	Tokenize(input string) ([]string, error)
}
