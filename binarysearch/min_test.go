package binarysearch_test

import "testing"

func TestGetMin(t *testing.T) {
	// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。

	// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

	// 请找出其中最小的元素。

	// 你可以假设数组中不存在重复元素。

	// 示例 1:
	// 输入: [3,4,5,1,2]
	// 输出: 1
	t.Log(getMin([]int{3, 4, 5, 1, 2}))

	// 示例 2:
	// 输入: [4,5,6,7,0,1,2]
	// 输出: 0
	t.Log(getMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func getMin(nums []int) int {
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

		if nums[midIndex] >= nums[l] {
			l = midIndex
		} else {
			r = midIndex
		}
	}

	return nums[midIndex]
}
