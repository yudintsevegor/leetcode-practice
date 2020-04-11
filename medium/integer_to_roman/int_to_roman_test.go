package integer_to_roman

import "testing"

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
