package arrays_test

import (
	"testing"
)

func TestGetMaxSortedDistance(t *testing.T) {
	// 获取无序数组中最大相邻差

	// [2, 6, 3, 4, 5, 10, 9]
	// 9 - 6 = 3
	t.Log(getMaxSortedDistance([]int{2, 6, 3, 4, 5, 10, 9}))
}

// 获取无序数组中最大相邻差, 计数排序思想
func getMaxSortedDistance(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	// 找到最大值和最小值
	max, min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			continue
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	// 偏移量
	d := min

	// 计数排序法
	numsCount := make([]int, max-min+1)
	for i := 0; i < len(nums); i++ {
		numsCount[nums[i]-d] = 1
	}

	// 找到相邻值为0的最大值
	var count, res int
	for i := 0; i < len(numsCount); i++ {
		if numsCount[i] == 0 {
			count++
		} else {
			count = 0
		}

		if count > res {
			res = count
		}
	}

	return res + 1
}
