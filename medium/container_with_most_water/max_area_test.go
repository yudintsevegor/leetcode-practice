package container_with_most_water

import (
	"reflect"
	"testing"
)

type Case struct {
	Input    []int
	Expected int
}

func TestFunc(t *testing.T) {
	for i, c := range makeCases() {
		result := maxArea(c.Input)
		if !reflect.DeepEqual(result, c.Expected) {
			t.Fatalf("[%d] Expected: %v; Result: %v", i, c.Expected, result)
		}
	}
}

func makeCases() []Case {
	return []Case{
		{
			Input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			Expected: 49,
		},
		{
			Input:    []int{1, 2, 3, 4},
			Expected: 4,
		},
	}
}
