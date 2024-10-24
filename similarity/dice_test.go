package similarity_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/similarity"
)

func TestDice(t *testing.T) {
	a := assert.New(t)
	d := similarity.Dice{}

	score, err := d.Compare("", "")
	a.NilNow(err)
	a.FloatEqual(score, 1.0, epsilon)

	score, err = d.Compare("Test String1", "Test String2")
	a.NilNow(err)
	a.FloatEqual(score, 0.5, epsilon)
}
