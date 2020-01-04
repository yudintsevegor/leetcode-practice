package findMedianSortedArrays

import "testing"

type Case struct {
	InputSliceOne  []int
	InputSliceTwo  []int
	ExpectedResult float64
}

func TestFunc(t *testing.T) {
	// t.Skip()
	cases := makeCases()
	for i, c := range cases {
		result := findMedianSortedArrays(c.InputSliceOne, c.InputSliceTwo)
		if result != c.ExpectedResult {
			t.Fatalf("[%d] Failed result for slices. Result: %v; Expected: %v",
				i, result, c.ExpectedResult)
		}
	}
}

func makeCases() []Case {
	return []Case{
		{
			InputSliceOne:  []int{1, 3},
			InputSliceTwo:  []int{2},
			ExpectedResult: 2.0,
		},
		{
			InputSliceOne:  []int{1, 2},
			InputSliceTwo:  []int{3, 4},
			ExpectedResult: 2.5,
		},
		{
			InputSliceOne:  []int{1},
			InputSliceTwo:  []int{4},
			ExpectedResult: 2.5,
		},
		{
			InputSliceOne:  []int{5, 7, 9},
			InputSliceTwo:  []int{4, 10},
			ExpectedResult: 7,
		},
		{
			InputSliceOne:  []int{5, 7, 7, 9},
			InputSliceTwo:  []int{4, 10, 10},
			ExpectedResult: 7,
		},
		{
			InputSliceOne:  []int{0, 0},
			InputSliceTwo:  []int{0, 0},
			ExpectedResult: 0,
		},
	}
}
