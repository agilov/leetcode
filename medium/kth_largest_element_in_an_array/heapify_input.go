package kth_largest_element_in_an_array

func findKthLargestViaHeapifyInput(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, len(nums), i)
	}

	for size := len(nums); k > 1; k-- {
		nums[0], nums[size-1] = nums[size-1], nums[0]
		size--
		heapify(nums, size, 0)
	}

	return nums[0]
}

func heapify(arr []int, size int, head int) {
	largest := head
	left := head*2 + 1
	right := head*2 + 2

	if left < size && arr[left] > arr[largest] {
		largest = left
	}

	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	if largest != head {
		arr[largest], arr[head] = arr[head], arr[largest]
		heapify(arr, size, largest)
	}
}
