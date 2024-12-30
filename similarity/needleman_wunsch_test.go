package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestNeedlemanWunsch(t *testing.T) {
	a := assert.New(t)
	mc := similarity.NeedlemanWunsch{}

	score, err := mc.Compare("", "")
	a.NilNow(err)
	a.FloatEqualNow(score, 1.0, epsilon)

	score, err = mc.Compare("Test String", "")
	a.NilNow(err)
	a.FloatEqual(score, 0.5, epsilon)

	score, err = mc.Compare("", "Test String")
	a.NilNow(err)
	a.FloatEqual(score, 0.5, epsilon)

	score, err = mc.Compare("Test String1", "Test String2")
	a.NilNow(err)
	a.FloatEqual(score, 23.0/24.0, epsilon)
}