package fizzbuzz

func fizzBuzz(num int) []string {
	result := make([]string, num+1)

	for i := 1; i <= num; i++ {
		if i%15 == 0 {
			result[i] = "FizzBuzz"
		} else if i%3 == 0 {
			result[i] = "Fizz"
		} else if i%5 == 0 {
			result[i] = "Buzz"
		} else {
			result[i] = int2string(i)
		}
	}

	return result[1:]
}

func int2string(num int) string {
	mapping := [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	if num < 10 {
		return string([]byte{mapping[num]})
	}

	result := []byte{}

	for ; num > 0; num /= 10 {
		result = append([]byte{mapping[num%10]}, result...)
	}

	return string(result)
}
