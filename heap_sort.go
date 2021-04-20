package gosort

func HeapSort(nums []int) {
	buildMaxHeap(nums)
	length := len(nums)
	for i := length - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, 0, i)
	}
}

func buildMaxHeap(nums []int) {
	length := len(nums)
	for i := length/2 - 1; i >= 0; i-- {
		heapify(nums, i, length)
	}
}

// 自i节点往下调整
func heapify(nums []int, i int, length int) {
	largest := i

	if 2*i+1 < length {
		if nums[i] < nums[2*i+1] {
			largest = 2*i + 1
		}
	}

	if 2*i+2 < length {
		if nums[largest] < nums[2*i+2] {
			largest = 2*i + 2
		}
	}

	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, largest, length)
	}
}
