package main

import (
	"strings"
)

func romanToInt(s string) int {
	dict := []struct {
		num string
		val int
	}{
		{"IV", 4},
		{"IX", 9},
		{"XL", 40},
		{"XC", 90},
		{"CD", 400},
		{"CM", 900},
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
	}

	result := 0

	for _, e := range dict {
		result += strings.Count(s, e.num) * e.val
		s = strings.ReplaceAll(s, e.num, "")
	}

	return result
}
