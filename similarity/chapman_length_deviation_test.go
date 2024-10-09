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
	a.Equal(1.0, score)

	score, err = cld.Compare("Test String1", "Test String2")
	a.NilNow(err)
	a.Equal(1.0, score)

	score, err = cld.Compare("ABC", "ABCDEF")
	a.NilNow(err)
	a.Equal(0.5, score)

	score, err = cld.Compare("ABCDEF", "ABC")
	a.NilNow(err)
	a.Equal(0.5, score)
}
