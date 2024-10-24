package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestCosine(t *testing.T) {
	a := assert.New(t)
	cs := similarity.Cosine{}

	score, err := cs.Compare("", "")
	a.NilNow(err)
	a.FloatEqual(score, 1.0, epsilon)

	score, err = cs.Compare("Test String1", "Test String2")
	a.NilNow(err)
	a.FloatEqual(score, 0.5, epsilon)
}
