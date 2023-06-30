package single_row_keyboard

import "testing"

// go test -v -run=TestCalculateTime ./easy/single_row_keyboard/*.go
func TestCalculateTime(t *testing.T) {
	result := calculateTime("abcdefghijklmnopqrstuvwxyz", "cba")

	if result != 4 {
		t.Errorf("Should be 4 got %v", result)
	}
}

// go test -v -run=_ -bench=BenchmarkCalculateTime ./easy/single_row_keyboard/*.go
func BenchmarkCalculateTime(b *testing.B) {
	kb := "abcdefghijklmnopqrstuvwxyz"
	word := "cba"
	for i := 0; i < b.N; i++ {
		calculateTime(kb, word)
	}
}
