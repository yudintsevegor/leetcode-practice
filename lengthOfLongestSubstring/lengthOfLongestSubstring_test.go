package lengthOfLongestSubstring

import "testing"

type Case struct {
	InputString    string
	ExpectedResult int
}

func TestFunc(t *testing.T) {
	// t.Skip()
	cases := makeCases()
	for i, c := range cases {
		result := lengthOfLongestSubstring(c.InputString)
		if result != c.ExpectedResult {
			t.Fatalf("[%d] Failed result for string: '%v'. Result: %v; Expected: %v",
				i, c.InputString, result, c.ExpectedResult)
		}
	}
}

func makeCases() []Case {
	return []Case{
		{
			InputString:    "abcabcbb",
			ExpectedResult: 3,
		},
		{
			InputString:    "bbbbb",
			ExpectedResult: 1,
		},
		{
			InputString:    "pwwkew",
			ExpectedResult: 3,
		},
		{
			InputString:    " ",
			ExpectedResult: 1,
		},
		{
			InputString:    "c",
			ExpectedResult: 1,
		},
		{
			InputString:    "ca",
			ExpectedResult: 2,
		},
		{
			InputString:    "cc",
			ExpectedResult: 1,
		},
		{
			InputString:    "dvdf",
			ExpectedResult: 3,
		},
		{
			InputString:    "rpdivsinokxnkctsfuk",
			ExpectedResult: 8,
		},
		{
			InputString:    "rpdivsinokxnkctsfukinavkn",
			ExpectedResult: 10,
		},
	}
}
