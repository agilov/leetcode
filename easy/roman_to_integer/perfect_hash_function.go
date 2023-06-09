package roman_to_integer

func getVal(s byte) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func RomanToIntPerfectHash(s string) int {
	var result, previous, current int

	for i := len(s) - 1; i >= 0; i-- {
		current = getVal(s[i])

		if current < previous {
			result -= current
		} else {
			result += current
		}

		previous = current
	}

	return result
}
