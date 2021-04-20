package gosort

func QuickSort(nums []int) []int {
	return _quickSort(nums, 0, len(nums))
}

func _quickSort(nums []int, start int, end int) []int {
	if start < end {
		index := partition(nums, start, end)
		_quickSort(nums, start, index-1)
		_quickSort(nums, index+1, end)
	}
	return nums
}

func partition(nums []int, start int, end int) int {
	pivot, index := start, start+1

	for i := start + 1; i < end; i++ {
		if nums[i] <= nums[pivot] {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	nums[index-1], nums[pivot] = nums[pivot], nums[index-1]

	return index - 1
}