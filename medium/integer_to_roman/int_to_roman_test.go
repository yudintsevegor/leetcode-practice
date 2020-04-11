package integer_to_roman

import (
	"reflect"
	"testing"
)

type Case struct {
	Input          int
	ExpectedResult string
}

func TestFunc(t *testing.T) {
	// t.Skip()
	for i, c := range makeCases() {
		result := intToRoman(c.Input)
		if result != c.ExpectedResult {
			t.Fatalf("[%d] Failed result for string: '%v'. Result: %v; Expected: %v",
				i, c.Input, result, c.ExpectedResult)
		}
	}
}

func makeCases() []Case {
	return []Case{
		{
			Input:          3,
			ExpectedResult: "III",
		},
		{
			Input:          4,
			ExpectedResult: "IV",
		},
		{
			Input:          9,
			ExpectedResult: "IX",
		},
		{
			Input:          1994,
			ExpectedResult: "MCMXCIV",
		},
	}
}

func TestSplitNum(t *testing.T) {
	t.Skip()
	for i, c := range makeSplitCases() {
		result := splitNumForOrderWithLoop(c.Input)
		if !reflect.DeepEqual(result, c.ExpectedResult) {
			t.Fatalf("[%d] Failed result for string: '%v'. Result: %v; Expected: %v",
				i, c.Input, result, c.ExpectedResult)
		}
	}
}

type SplitCase struct {
	Input          int
	ExpectedResult []int
}

func makeSplitCases() []SplitCase {
	return []SplitCase{
		{
			Input:          3,
			ExpectedResult: []int{3},
		},
		{
			Input:          40,
			ExpectedResult: []int{0, 4},
		},
		{
			Input:          999,
			ExpectedResult: []int{9, 9, 9},
		},
		{
			Input:          1994,
			ExpectedResult: []int{4, 9, 9, 1},
		},
	}
}
