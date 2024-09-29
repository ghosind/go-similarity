package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestDice(t *testing.T) {
	a := assert.New(t)
	d := similarity.Dice{}

	a.True(isSameFloat64(1.0, d.Compare("", "")))
	a.True(isSameFloat64(0.5, d.Compare("Test String1", "Test String2")))
}
