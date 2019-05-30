package binarysearch_test

import "testing"

func TestBSearch3(t *testing.T) {
	// 查找第一个大于等于给定值的元素
	nums := []int{1, 2, 3, 3, 4, 4, 5, 6, 7, 7, 7, 8, 8, 9}

	// 2
	t.Log(bsearch3(nums, 3))

	// 4
	t.Log(bsearch3(nums, 4))

	// 8
	t.Log(bsearch3(nums, 7))

	// 11
	t.Log(bsearch3(nums, 8))
}

// 查找第一个大于等于给定值的元素
func bsearch3(nums []int, n int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>1

		if nums[mid] >= n {
			if mid == 0 || nums[mid-1] < n {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}

	return -1
}
