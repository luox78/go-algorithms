package binarysearch_test

import "testing"

func TestFindMin(t *testing.T) {
	// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。

	// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

	// 请找出其中最小的元素。

	// 注意数组中可能存在重复的元素。

	// 示例 1：
	// 输入: [1,3,5]
	// 输出: 1
	t.Log(findMin([]int{1, 3, 5}))

	// 示例 2：
	// 输入: [2,2,2,0,1]
	// 输出: 0
	t.Log(findMin([]int{2, 2, 2, 0, 1}))
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	midIndex := l

	for l != r && nums[l] >= nums[r] {
		if r-l == 1 {
			midIndex = r
			break
		}

		midIndex = l + (r-l)>>1

		if nums[l] == nums[r] && nums[midIndex] == nums[l] {
			res := nums[l]
			for i := l + 1; i <= r; i++ {
				if res > nums[i] {
					res = nums[i]
				}
			}
			return res
		}

		if nums[midIndex] >= nums[l] {
			l = midIndex
		} else {
			r = midIndex
		}
	}

	return nums[midIndex]
}
