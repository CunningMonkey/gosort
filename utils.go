package gosort

func MaxValue(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

func MinAndMaxValue(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	max, min := nums[0], nums[0]

	for _, num := range nums {
		if max < num {
			max = num
		}
		if min > num {
			min = num
		}
	}

	return min, max
}

func MaxBits(nums []int) int {
	maxBits := 1
	bits := 1
	for _, num := range nums {
		bits = IntegerBits(num)
		if maxBits < bits {
			maxBits = bits
		}
	}

	return maxBits
}

func IntegerBits(num int) int {
	bits := 1
	for num > 0 {
		num /= 10
		bits++
	}
	return bits-1
}

func GetBit(num int, location int) int {
	for location > 1 {
		num /= 10
		location--
	}
	return num % 10
}
