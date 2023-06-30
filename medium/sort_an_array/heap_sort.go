package sort_an_array

func heapsortArray(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, len(nums), i)
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}

	return nums
}

func heapify(arr []int, size int, head int) {
	left := head*2 + 1
	right := head*2 + 2
	largest := head

	if left < size && arr[left] > arr[largest] {
		largest = left
	}

	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	if head != largest {
		arr[head], arr[largest] = arr[largest], arr[head]
		heapify(arr, size, largest)
	}
}
