package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestOverlapCoefficient(t *testing.T) {
	a := assert.New(t)
	oc := similarity.OverlapCoefficient{}

	a.True(isSameFloat64(1.0, oc.Compare("", "")))
	a.True(isSameFloat64(0.5, oc.Compare("Test String1", "Test String2")))
}
