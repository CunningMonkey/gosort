package gosort

func Merge_sort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}

	middle := length / 2
	nums1, nums2 := nums[0:middle], nums[middle:length]
	return merge(Merge_sort(nums1), Merge_sort(nums2))
}

func merge(nums1 []int, nums2 []int) (nums []int) {
	index, index1, index2 := 0, 0, 0
	nums = make([]int, len(nums1)+len(nums2))
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] <= nums2[index2] {
			nums[index] = nums1[index1]
			index++
			index1++
		} else {
			nums[index] = nums2[index2]
			index++
			index2++
		}
	}

	for index1 < len(nums1) {
		nums[index] = nums1[index1]
		index++
		index1++
	}

	for index2 < len(nums2) {
		nums[index] = nums2[index2]
		index++
		index2++
	}

	return nums
}
