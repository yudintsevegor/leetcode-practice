package stringsToInteger

import "testing"

type Case struct {
	Input    string
	Expected int
}

func TestFunc(t *testing.T) {
	for i, c := range getCases() {
		result := myAtoi(c.Input)
		if result != c.Expected {
			t.Fatalf("[%d] Expected: %v; Result: %v; Input: %v", i, c.Expected, result, c.Input)
		}
	}
}

func getCases() []Case {
	return []Case{
		{
			Input:    "42",
			Expected: 42,
		},
		{
			Input:    "   -42",
			Expected: -42,
		},
		{
			Input:    "4193 with words",
			Expected: 4193,
		},
		{
			Input:    "words and 987",
			Expected: 0,
		},
		{
			Input:    "-91283472332",
			Expected: -2147483648,
		},
		{
			Input:    "2147483648",
			Expected: 2147483647,
		},
		{
			Input:    "20000000000000000000",
			Expected: 2147483647,
		},
		{
			Input:    "-5-",
			Expected: -5,
		},
		{
			Input:    "123-",
			Expected: 123,
		},
	}
}
