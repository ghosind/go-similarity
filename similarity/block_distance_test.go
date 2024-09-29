package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestBlockDistance(t *testing.T) {
	a := assert.New(t)
	bd := similarity.BlockDistance{}

	a.Equal(0.5, bd.Compare("Test String1", "Test String2"))
	a.Equal(1.0, bd.Compare("", ""))
}
