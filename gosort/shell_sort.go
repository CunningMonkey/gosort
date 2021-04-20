package gosort

func ShellSort(nums []int) {
	length := len(nums)
	gap := length / 2

	for gap > 0 {
		for i := gap; i < length; i++ {
			for j := i; j >= gap; j -= gap {
				if nums[j] < nums[j-gap] {
					nums[j], nums[j-gap] = nums[j-gap], nums[j]
				}
			}
		}
		gap /= 2
	}
}

// func main() {
// 	nums := []int{99, 5, 69, 33, 56, 13, 22, 55, 77, 48, 12, 88, 2, 69, 99}
// 	shellSort(nums)
// 	fmt.Printf("%v", nums)
// }
