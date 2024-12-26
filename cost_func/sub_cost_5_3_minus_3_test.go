package cost_func_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-similarity/cost_func"
)

func TestSubCost5_3Minus3(t *testing.T) {
	a := assert.New(t)
	f := new(cost_func.SubCost5_3Minus3)

	s1 := "hello world AAAAAAA BBB ABCDEF this is a test"
	s2 := "jello wrd AAAAAAA BBB ABCDEF this is a test"
	a.Equal(f.Cost(s1, 0, s2, 0), -3.0)
	a.Equal(f.Cost(s1, 2, s2, 2), 5.0)
	a.Equal(f.Cost(s1, 7, s2, 7), -3.0)
	a.Equal(f.Cost(s1, 10, s2, 10), -3.0)
	a.Equal(f.Cost(s1, 22, s2, 3), -3.0)
	a.Equal(f.Cost(s1, 1, s2, 4), 3.0)
	a.Equal(f.Cost(s1, 50, s2, 0), -3.0)
}
