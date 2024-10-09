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
	a.Equal(0.5, score)

	score, err = bd.Compare("", "")
	a.NilNow(err)
	a.Equal(1.0, score)
}
