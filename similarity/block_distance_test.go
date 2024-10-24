package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestBlockDistance(t *testing.T) {
	a := assert.New(t)
	bd := similarity.BlockDistance{}

	score, err := bd.Compare("Test String1", "Test String2")
	a.NilNow(err)
	a.FloatEqual(score, 0.5, epsilon)

	score, err = bd.Compare("", "")
	a.NilNow(err)
	a.FloatEqual(score, 1.0, epsilon)
}
