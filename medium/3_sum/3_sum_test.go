package __sum

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	// t.Skip()
	for i, c := range makeCases() {
		result := threeSum(c.Input)
		if !reflect.DeepEqual(result, c.ExpectedResult) {
			t.Fatalf("[%d] result for input: %v. Result: %v; Expected: %v", i, c.Input, result, c.ExpectedResult)
		}
	}
}

type Case struct {
	Input          []int
	ExpectedResult [][]int
}

func makeCases() []Case {
	return []Case{
		{
			Input: []int{-1, 0, 1, 2, -1, 4},
			ExpectedResult: [][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			},
		},
		{
			Input: []int{0, 0, 0},
			ExpectedResult: [][]int{
				{0, 0, 0},
			},
		},
		{
			Input: []int{1, -1, -1, 0},
			ExpectedResult: [][]int{
				{-1, 0, 1},
			},
		},
	}
}
