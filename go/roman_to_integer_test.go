package main

import "testing"

// go test -v ./roman_to_integer.go ./roman_to_integer_test.go
func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		roman  string
		expect int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MDLVI", 1556},
		{"MCMXCI", 1991},
		{"MMMCMXCIX", 3999},
		{"MMMCMXCVIII", 3998},
	}

	for _, tc := range testCases {
		got := romanToInt(tc.roman)

		if got != tc.expect {
			t.Errorf("romanToInt(\"%s\"): expected %d, but got %d", tc.roman, tc.expect, got)
		}
	}
}

// go test -bench=. ./roman_to_integer.go ./roman_to_integer_test.go
func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		romanToInt("CMIII")
	}
}
