package single_row_keyboard

func calculateTime(keyboard string, word string) int {
	mapping := [26]int{}

	for i, c := range []byte(keyboard) {
		mapping[c-'a'] = i
	}

	result := 0
	cur := 0
	var abs = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	for _, c := range []byte(word) {
		idx := mapping[c-'a']
		result += abs(idx - cur)
		cur = idx
	}

	return result
}
