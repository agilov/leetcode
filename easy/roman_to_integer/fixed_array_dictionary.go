package roman_to_integer

var roman = [256]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func RomanToIntFixedArray(s string) int {

	var result, previous, current int

	for i := len(s) - 1; i >= 0; i-- {
		current = roman[s[i]]

		if roman[s[i]] < previous {
			result -= current
		} else {
			result += current
		}

		previous = current
	}

	return result
}
