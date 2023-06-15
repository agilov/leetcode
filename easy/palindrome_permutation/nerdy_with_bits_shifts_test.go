package palindrome_permutation

import "testing"

// go test -v -run=TestCanPermutePalindrome ./easy/palindrome_permutation/*.go
func TestCanPermutePalindrome(t *testing.T) {
	cases := map[string]bool{
		"code": false,
		"aab":  true,
		"aaa":  true,
		"jop":  false,
		"kek":  true,
		"":     true,
		"a":    true,
	}

	for input, exect := range cases {
		ans := canPermutePalindrome(input)

		if ans != exect {
			t.Errorf("For string %v expected %v but got %v", input, exect, ans)
		}
	}
}
