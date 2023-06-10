package roman_to_integer

func RomanToIntFixedArray(s string) int {
	var roman = [256]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	var result, previous, current int

	for i := len(s) - 1; i >= 0; i-- {
		current = roman[s[i]]

		if current < previous {
			result -= current
		} else {
			result += current
		}

		previous = current
	}

	return result
}
