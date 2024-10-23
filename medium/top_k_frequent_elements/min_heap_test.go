package top_k_requent_elements

import (
    "testing"
	"math/rand/v2"
)

// go test -v -run=TestTopKFrequentElements ./medium/top_k_frequent_elements/*.go
func TestTopKFrequentElements(t *testing.T) {
    testCases := []struct {
        n []int
        k int
        e []int
    }{
        {[]int{}, 0, []int{}},
        {[]int{1,1,2}, 1, []int{1}},
        {[]int{1,1,2,2,2,3}, 2, []int{2,1}},
        {[]int{1,1,2,2,2,3}, 3, []int{2,1,3}},
        {[]int{1,1,2,2,2,3,4,4,4,4,4}, 3, []int{2,1,4}},
        {[]int{1,2,3,4,5,6,1,1,1,1,1,2,2,2,2,2,3,3,3,3,4,4,4,5,5,5,6,6,123,435,141,89,99}, 6, []int{1,2,3,4,5,6}},
        {[]int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 4, 6, 6, 6, 6, 6, 6, 6, 6, 6, 9, 1, 4}, 3, []int{6, 4, 1}},
    }
	for _, c := range testCases {
		rand.Shuffle(len(c.n), func(i,j int) {
			c.n[i], c.n[j] = c.n[j], c.n[i]
		})
		result := topKFrequent(c.n, c.k)
		m := map[int]bool{}
		for _, v := range c.e {
			m[v] = true
		}
		if len(result) != len(c.e) {
			t.Fatalf("Expected result len %d but got %d", len(c.e), len(result))
		}
		
		for _, e := range result {
			if !m[e] {
				t.Fatalf("%d should not be present in result", e)
			}
		}
	}
}
