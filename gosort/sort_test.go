package gosort

import (
	"testing"
)

func TestRadixSort(t *testing.T) {
	nums := []int{99, 5, 69, 33, 56, 13, 22, 55, 77, 48, 12, 88, 2, 69, 99}
	resTest := []int{2, 5, 12, 13, 22, 33, 48, 55, 56, 69, 69, 77, 88, 99, 99}
	RadixSort(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] != resTest[i] {
			t.Fail()
		}
	}
}
