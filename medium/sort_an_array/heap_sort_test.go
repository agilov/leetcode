package sort_an_array

import "testing"

func assertSlicesEqual(input, result, expect []int, t *testing.T) {
	for i, v := range result {
		if expect[i] != v {
			t.Errorf("For input array %v expected sorted result %v but got %v", input, result, expect)
			return
		}
	}
}

// go test -v -run=TestMergeSort ./medium/sort_an_array/*.go
func TestMergeSort(t *testing.T) {
	cases := [...][2][]int{
		[2][]int{
			[]int{1, 2, 4, 5, 3},
			[]int{1, 2, 3, 4, 5},
		},
		[2][]int{
			[]int{2, 4, 5, 3, -100},
			[]int{-100, 2, 3, 4, 5},
		},
		[2][]int{
			[]int{12, 2134, 1234, 89090, 123, 123, 2, 789, 9, 0, 4, 0, 3},
			[]int{0, 0, 2, 3, 4, 9, 12, 123, 123, 789, 1234, 2134, 89090},
		},
	}

	for _, c := range cases {
		assertSlicesEqual(c[0], heapsortArray(c[0]), c[1], t)
	}
}

// go test -v -bench=BenchmarkMergeSort ./medium/sort_an_array/*.go
func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		heapsortArray([]int{12, 2134, 1234, 19, 12, 2134, 1234, 89090, 123, 123, 2, 5, 63, 789, 9, 0, 4, 0, 3, 123, 123, 2, 789, 9, 0, 4, 0, 3})
	}
}
