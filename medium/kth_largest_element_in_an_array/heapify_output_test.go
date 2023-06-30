package kth_largest_element_in_an_array

import "testing"

// go test -v -run=TestFindKthLargestViaHeapifyOutput ./medium/kth_largest_element_in_an_array/*.go
func TestFindKthLargestViaHeapifyOutput(t *testing.T) {
	cases := [...]struct {
		nums   []int
		k      int
		expect int
	}{
		{[]int{1, 2, 3}, 1, 3},
		{[]int{8, 5, 6, 3, 1, 0, -1}, 3, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{3, 2, 3, 8, 12, 1, 2, 11, 84, 4, 5, 5, 6}, 4, 8},
		{[]int{3, 99, 2, 73, 3, 8, 12, 1, 2, 11, 84, 4, 5, 65, 5, 6, 101, 102}, 10, 6},
	}

	for _, c := range cases {
		result := findKthLargestViaHeapifyOutput(c.nums, c.k)
		if result != c.expect {
			t.Errorf("For input nums %v expected %v-th largest element is %v but got %v", c.nums, c.k, c.expect, result)
		}
	}
}

// go test -v -run=^$ -bench=BenchmarkFindKthLargestViaHeapifyOutput ./medium/kth_largest_element_in_an_array/*.go
func BenchmarkFindKthLargestViaHeapifyOutput(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findKthLargestViaHeapifyOutput([]int{3, 5, 6, 83, -138, 12, 12123, 94923, 56, 82, 2, 3, 1, 2, 4, 3, -13, 1, 89, 10, -123, 32, 60, 12, 1, 1, 89, 10, -112, 12123, 94923, 56, 82, 2, 323, 32, 60, 12, 12123, 94923, 56, 82, 2, 3, 1, 2, 4, 5, 5, 6}, 20)
	}
}
