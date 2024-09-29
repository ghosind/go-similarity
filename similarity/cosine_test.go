package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestCosine(t *testing.T) {
	a := assert.New(t)
	cs := similarity.Cosine{}

	a.True(isSameFloat64(1.0, cs.Compare("", "")))
	a.True(isSameFloat64(0.5, cs.Compare("Test String1", "Test String2")))
}
