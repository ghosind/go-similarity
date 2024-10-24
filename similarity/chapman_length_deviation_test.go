package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestChapmanLengthDeviation(t *testing.T) {
	a := assert.New(t)
	cld := similarity.ChapmanLengthDeviation{}

	score, err := cld.Compare("", "")
	a.NilNow(err)
	a.FloatEqual(score, 1.0, epsilon)

	score, err = cld.Compare("Test String1", "Test String2")
	a.NilNow(err)
	a.FloatEqual(score, 1.0, epsilon)

	score, err = cld.Compare("ABC", "ABCDEF")
	a.NilNow(err)
	a.FloatEqual(score, 0.5, epsilon)

	score, err = cld.Compare("ABCDEF", "ABC")
	a.NilNow(err)
	a.FloatEqual(score, 0.5, epsilon)
}
