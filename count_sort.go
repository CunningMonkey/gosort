package gosort

func CountSort(nums []int) {
	maxValue := MaxValue(nums)
	buckets := make([]int, maxValue+1)

	for _, num := range nums {
		buckets[num] += 1
	}

	index := 0
	for k, bucket := range buckets {
		for bucket > 0 {
			nums[index] = k
			index++
			bucket--
		}
	}
}
