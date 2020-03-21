package reverse_integer

import (
	"testing"
)

type Case struct {
	Input    int
	Expected int
}

func TestFunc(t *testing.T) {
	for i, c := range getCases() {
		result := reverse(c.Input)
		if result != c.Expected {
			t.Fatalf("[%d] Expected: %v; Result: %v; Input: %v", i, c.Expected, result, c.Input)
		}
	}
}

func getCases() []Case {
	return []Case{
		{
			Input:    123,
			Expected: 321,
		},
		{
			Input:    -123,
			Expected: -321,
		},
		{
			Input:    120,
			Expected: 21,
		},
		{
			Input:    100,
			Expected: 1,
		},
		{
			Input:    1,
			Expected: 1,
		},
		{
			Input:    10,
			Expected: 1,
		},
		{
			Input:    1534236469,
			Expected: 0,
		},
	}
}
