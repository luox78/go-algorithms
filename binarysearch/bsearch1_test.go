package binarysearch_test

import "testing"

func TestBSearch1(t *testing.T) {
	// 查找第一个值等于给定值的元素
	nums := []int{1, 2, 3, 3, 4, 4, 5, 6, 7, 7, 7, 8, 8, 9}

	// 2
	t.Log(bsearch1(nums, 3))

	// 4
	t.Log(bsearch1(nums, 4))

	// 8
	t.Log(bsearch1(nums, 7))
}

// 查找第一个值等于给定值的元素
func bsearch1(nums []int, n int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>1

		if nums[mid] == n {
			// 判断前一个值是否等于给定值
			if mid == 0 || nums[mid-1] != n {
				return mid
			} else {
				high = mid - 1
			}
		} else if nums[mid] > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
