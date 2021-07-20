package main

import "testing"

// run-length encode a string
var encodeTests = []struct {
	input    string
	expected string
}{
	{
		input:    "",
		expected: "",
	},
	{
		input:    "XYZ",
		expected: "XYZ",
	},
	{
		input:    "AABBBCCCC",
		expected: "2A3B4C",
	},
	{
		input:    "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB",
		expected: "12WB12W3B24WB",
	},
	{
		input:    "  hsqq qww  ",
		expected: "2 hs2q q2w2 ",
	},
	{
		input:    "한한한AAA",
		expected: "3한3A",
	},
}

func TestRunLengthEncode(t *testing.T) {
	for _, test := range encodeTests {
		if actual := RunLengthEncode(test.input); actual != test.expected {
			t.Errorf("FAIL- RunLengthEncode(%s) = %q, expected %q.",
				test.input, actual, test.expected)
		}
		t.Logf("PASS RunLengthEncode")
	}
}
