package gosort

import "testing"

func TestIntegerBits(t *testing.T) {
	nums := []int{1, 2, 10, 100, 9999, 10001}
	resTest := []int{1, 1, 2, 3, 4, 5}
	for i, num := range nums {
		res := IntegerBits(num)
		if resTest[i] != res {
			t.Fatalf("IntegerBits Faild! true : %d, false :%d", resTest[i], res)
		}
	}
}

func TestMaxBits(t *testing.T) {
	nums := []int{1, 2, 10, 100, 9999, 10001}
	res := MaxBits(nums)
	if res != 5 {
		t.Fatalf("MaxBits Fail! %d", res)
	}
}

func TestGetBit(t *testing.T) {
	num := 120301312
	resTest := []int {2, 1, 3, 1, 0, 3, 0, 2, 1}
	for i := 0; i < len(resTest); i++ {
		res := GetBit(num, i+1)
		if res != resTest[i] {
			t.Fatalf("GetBit Fail! index: %d true: %d, false: %d", i, resTest[i], res)
		}
	}
}