package zig_zag_convert

import "testing"

type Case struct {
	InputString string
	NumRows     int
	Expected    string
}

func TestFunc(t *testing.T) {
	for i, c := range getCases() {
		converted := convert(c.InputString, c.NumRows)
		if converted != c.Expected {
			t.Fatalf("[%v] Rows: %v; Result: %v; Expected: %v", i, c.NumRows, converted, c.Expected)
		}
	}
}

func getCases() []Case {
	return []Case{

		{
			InputString: "PAYPALISHIRING",
			NumRows:     3,
			Expected:    "PAHNAPLSIIGYIR",
		},
		{
			InputString: "PAYPALISHIRING",
			NumRows:     4,
			Expected:    "PINALSIGYAHRPI",
		},
		{
			InputString: "PAYPALISHIRING",
			NumRows:     1,
			Expected:    "PAYPALISHIRING",
		},
		{
			InputString: "PAYPALISHIRING",
			NumRows:     2,
			Expected:    "PYAIHRNAPLSIIG",
		},
	}
}
