package binarysearch_test

import "testing"

func TestBSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 3
	t.Log(bsearch(nums, 4))
	t.Log(bsearchInternally(nums, 4, 0, 8))

	// 4
	t.Log(bsearch(nums, 5))
	t.Log(bsearchInternally(nums, 5, 0, 8))

	// 0
	t.Log(bsearch(nums, 1))
	t.Log(bsearchInternally(nums, 1, 0, 8))

	// 8
	t.Log(bsearch(nums, 9))
	t.Log(bsearchInternally(nums, 9, 0, 8))
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

// 简单二分查找递归实现
func bsearchInternally(nums []int, n, low, high int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)>>1

	if nums[mid] == n {
		return mid
	} else if nums[mid] > n {
		return bsearchInternally(nums, n, low, mid-1)
	} else {
		return bsearchInternally(nums, n, mid+1, high)
	}

}
