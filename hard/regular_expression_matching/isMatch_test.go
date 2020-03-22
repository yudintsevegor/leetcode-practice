package regular_expression_matching

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for i, c := range createMainCases() {
		result := isMatch(c.InputString, c.InputPattern)
		if result != c.Expected {
			t.Fatalf("[%d] InputString: %v; InputPattern: %v; Result: %v; Expected: %v;",
				i, c.InputString, c.InputPattern, result, c.Expected)
		}
	}
}

type MainCase struct {
	InputString  string
	InputPattern string
	Expected     bool
}

func createMainCases() []MainCase {
	return []MainCase{
		{
			InputString:  "aa",
			InputPattern: "a",
			Expected:     false,
		},
		{
			InputString:  "aa",
			InputPattern: "a*",
			Expected:     true,
		},
		{
			InputString:  "aab",
			InputPattern: "c*a*b",
			Expected:     true,
		},
		{
			InputString:  "mississippi",
			InputPattern: "mis*is*p*.",
			Expected:     false,
		},
		{
			InputString:  "ab",
			InputPattern: ".*",
			Expected:     true,
		},
		{
			InputString:  "ab",
			InputPattern: ".*c",
			Expected:     false,
		},
		{
			InputString:  "aaa",
			InputPattern: ".*",
			Expected:     true,
		},
		{
			InputString:  "aaa",
			InputPattern: "aaaa",
			Expected:     false,
		},
		{
			InputString:  "aaa",
			InputPattern: "a*a",
			Expected:     true,
		},
		{
			InputString:  "aaa",
			InputPattern: "ab*a*c*a",
			Expected:     true,
		},
		{
			InputString:  "a",
			InputPattern: "ab*",
			Expected:     true,
		},
		{
			InputString:  "ab",
			InputPattern: ".*..",
			Expected:     true,
		},
	}
}

func TestCutTail(t *testing.T) {
	t.Skip()
	for i, c := range createCutCases() {
		result := cutTail(c.Input)
		if !reflect.DeepEqual(result, c.Expected) {
			t.Fatalf("[%d] Result: %v; Expected: %v;", i, result, c.Expected)
		}
	}
}

type CutCase struct {
	Input    []string
	Expected []string
}

func createCutCases() []CutCase {
	return []CutCase{
		{
			Input:    []string{"a", "b", "c"},
			Expected: []string{"a", "b", "c"},
		},
		{
			Input:    []string{"*", "b", "c"},
			Expected: []string{"*", "b", "c"},
		},
		{
			Input:    []string{"*", "b", "*"},
			Expected: []string{"*"},
		},
		{
			Input:    []string{"*"},
			Expected: []string{"*"},
		},
		{
			Input:    []string{"*", "c", "*", "a"},
			Expected: []string{"*", "a"},
		},
	}
}
