package binarysearch_test

import "testing"

func TestBSearch2(t *testing.T) {
	// 查找最后一个值等于给定值的元素
	nums := []int{1, 2, 3, 3, 4, 4, 5, 6, 7, 7, 7, 8, 8, 9}

	// 3
	t.Log(bsearch2(nums, 3))

	// 5
	t.Log(bsearch2(nums, 4))

	// 10
	t.Log(bsearch2(nums, 7))

	// 12
	t.Log(bsearch2(nums, 8))
}

// 查找最后一个值等于给定值的元素
func bsearch2(nums []int, n int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>1

		if nums[mid] == n {
			if mid == len(nums)-1 || nums[mid+1] != n {
				return mid
			}
			low = mid + 1
		} else if nums[mid] > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
