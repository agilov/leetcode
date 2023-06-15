package palindrome_permutation

func canPermutePalindrome(s string) bool {
	bitset := 0

	for _, c := range []byte(s) {
		bitset ^= 1 << (c - 'a')
	}

	ones := 0
	for ; bitset > 0; bitset >>= 1 {
		ones += bitset & 1
	}

	return ones <= 1
}
