package reverseInteger

import "testing"

type Case struct {
	Input    int
	Expected int
}

func TestFunc(t *testing.T) {
}

func getCases() []Case {
	return []Case{
		{
			Input:    0,
			Expected: 0,
		},
	}
}
