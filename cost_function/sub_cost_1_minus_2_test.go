package cost_function_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/cost_function"
)

func TestSubCost1Minus2(t *testing.T) {
	a := assert.New(t)
	f := new(cost_function.SubCost1Minus2)

	s1 := "hello world AAAAAAA BBB ABCDEF this is a test"
	s2 := "jello wrd AAAAAAA BBB ABCDEF this is a test"
	a.Equal(f.Cost(s1, 0, s2, 0), -2.0)
	a.Equal(f.Cost(s1, 2, s2, 2), 1.0)
	a.Equal(f.Cost(s1, 7, s2, 7), -2.0)
	a.Equal(f.Cost(s1, 10, s2, 10), -2.0)
	a.Equal(f.Cost(s1, 22, s2, 3), -2.0)
	a.Equal(f.Cost(s1, 50, s2, 0), 0.0)
}
