package gosort

func RadixSort(nums []int) {
	maxBits := MaxBits(nums)
	buckets := make([][][]int, maxBits)
	for i := 0; i < maxBits; i++ {
		buckets[i] = make([][]int, 10)
		for _, num := range nums {
			bucket := buckets[i][GetBit(num, i+1)]
			buckets[i][GetBit(num, i + 1)] = append(bucket, num)
		}

		index := 0
		for k := 0; k < 10; k++ {
			for _, num := range buckets[i][k] {
				nums[index] = num
				index++
			}
		}
	}

	num := 1
	num++
}
