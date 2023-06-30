package kth_largest_element_in_an_array

func findKthLargestViaHeapifyOutput(nums []int, k int) int {
	for i := k/2 - 1; i >= 0; i-- {
		minHeapify(nums, k, i)
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			minHeapify(nums, k, 0)
		}
	}

	return nums[0]
}

func minHeapify(arr []int, size int, head int) {
	smallest := head
	left := head*2 + 1
	right := head*2 + 2

	if left < size && arr[left] < arr[smallest] {
		smallest = left
	}

	if right < size && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest != head {
		arr[smallest], arr[head] = arr[head], arr[smallest]
		minHeapify(arr, size, smallest)
	}
}
