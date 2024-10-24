package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestOverlapCoefficient(t *testing.T) {
	a := assert.New(t)
	oc := similarity.OverlapCoefficient{}

	score, err := oc.Compare("", "")
	a.NilNow(err)
	a.FloatEqual(score, 1.0, epsilon)

	score, err = oc.Compare("Test String1", "Test String2")
	a.NilNow(err)
	a.FloatEqual(score, 0.5, epsilon)
}
