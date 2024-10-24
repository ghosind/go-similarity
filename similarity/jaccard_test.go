package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestJaccard(t *testing.T) {
	a := assert.New(t)
	j := similarity.Jaccard{}

	score, err := j.Compare("", "")
	a.NilNow(err)
	a.FloatEqual(score, 1.0, epsilon)

	score, err = j.Compare("Test String1", "Test String2")
	a.NilNow(err)
	a.FloatEqual(score, 1.0/3.0, epsilon)
}
