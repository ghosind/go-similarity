package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestLevenshtein(t *testing.T) {
	a := assert.New(t)
	l := similarity.Levenshtein{}

	score, err := l.Compare("", "")
	a.NilNow(err)
	a.FloatEqual(score, 1.0, epsilon)

	score, err = l.Compare("Test String", "")
	a.NilNow(err)
	a.FloatEqual(score, 0.0, epsilon)

	score, err = l.Compare("Test String1", "Test String2")
	a.NilNow(err)
	a.FloatEqual(score, 11.0/12.0, epsilon)
}
