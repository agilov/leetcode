package fizzbuzz

import "testing"

func TestFizzBuzzNaive(t *testing.T) {
	result := fizzBuzz(33)

	if result[0] != "1" {
		t.Error("Shoul be 1")
	}

	if result[1] != "2" {
		t.Error("Shoul be 2")
	}

	if result[2] != "Fizz" {
		t.Error("Shoul be Fizz")
	}
}
