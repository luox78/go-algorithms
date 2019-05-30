package binarysearch_test

import "testing"

func TestBSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 3
	t.Log(bsearch(nums, 4))

	// 4
	t.Log(bsearch(nums, 5))

	// 0
	t.Log(bsearch(nums, 1))

	// 8
	t.Log(bsearch(nums, 9))
}

// 简单二分查找
func bsearch(nums []int, n int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>1

		if nums[mid] == n {
			return mid
		}

		if nums[mid] > n {
			// 在左区间
			high = mid - 1
		} else {
			// 在右区间
			low = mid + 1
		}
	}

	return -1
}
