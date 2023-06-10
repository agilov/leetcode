package roman_to_integer

import "testing"

// go test -v ./easy/roman_to_integer/fixed_array_dictionary*.go

// go test -v -run TestRomanToIntFixedArray ./easy/roman_to_integer/fixed_array_dictionary*.go
func TestRomanToIntFixedArray(t *testing.T) {
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
		got := RomanToIntFixedArray(tc.roman)

		if got != tc.expect {
			t.Errorf("romanToInt(\"%s\"): expected %d, but got %d", tc.roman, tc.expect, got)
		}
	}
}

// go test -bench=BenchmarkRomanToIntFixedArray ./easy/roman_to_integer/fixed_array_dictionary*.go
func BenchmarkRomanToIntFixedArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomanToIntFixedArray("MMMCMXCVIII")
		RomanToIntFixedArray("LVIII")
		RomanToIntFixedArray("IX")
		RomanToIntFixedArray("MCMXCI")
	}
}
