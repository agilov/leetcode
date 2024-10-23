package top_k_requent_elements

type minHeap [][]int

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	heap := make(minHeap, 0, k)
	
	for num, freq := range m {
		
		if len(heap) < cap(heap) {
			heap = append(heap, []int{freq, num})
			swim(heap, len(heap)-1)
		} else if heap[0][0] < freq {
			heap[0] = []int{freq, num}
			sink(heap, 0)
		}
	}
	result := []int{}
	for _, v := range heap {
		result = append(result, v[1])
	}
	return result
}

func sink(h minHeap, pos int) {
	max := len(h) - 1
	if pos == max {
		return
	}

	child1 := pos*2 + 1
	child2 := pos*2 + 2

	if child1 > max {
		return
	}
	swap := child1
	if child2 <= max && h[child2][0] < h[child1][0] {
		swap = child2
	}

	if h[pos][0] > h[swap][0] {
		h[pos], h[swap] = h[swap], h[pos]
		sink(h, swap)
	}
}

func swim(h minHeap, pos int) {
	if pos == 0 {
		return
	}

	parent := (pos - 1)/2
	if parent >= 0 && h[pos][0] < h[parent][0] {
		
		h[pos], h[parent] = h[parent], h[pos]
		swim(h, parent)
	}
}
