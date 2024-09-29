package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestEuclideanDistance(t *testing.T) {
	a := assert.New(t)
	ed := similarity.EuclideanDistance{}

	a.True(isSameFloat64(1.0, ed.Compare("", "")))
	a.True(isSameFloat64(0.5, ed.Compare("Test String1", "Test String2")))
}
