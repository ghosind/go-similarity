package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestChapmanLengthDeviation(t *testing.T) {
	a := assert.New(t)
	cld := similarity.ChapmanLengthDeviation{}

	a.Equal(1.0, cld.Compare("", ""))
	a.Equal(1.0, cld.Compare("Test String1", "Test String2"))
	a.Equal(0.5, cld.Compare("ABC", "ABCDEF"))
	a.Equal(0.5, cld.Compare("ABCDEF", "ABC"))
}
