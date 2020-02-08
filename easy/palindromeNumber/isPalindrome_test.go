package palindromeNumber

import "testing"

type Case struct {
	Input    int
	Expected bool
}

func TestFunc(t *testing.T) {
	for i, c := range getCases() {
		result := isPalindrome(c.Input)
		if result != c.Expected {
			t.Fatalf("[%d] Input: %v; Result: %v; Expected: %v;", i, c.Input, result, c.Expected, )
		}
	}
}

func getCases() []Case {
	return []Case{
		{
			Input:    121,
			Expected: true,
		},
		{
			Input:    -121,
			Expected: false,
		},
		{
			Input:    10,
			Expected: false,
		},
	}
}

