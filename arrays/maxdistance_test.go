package arrays_test

import (
	"testing"
)

func TestGetMaxSortedDistance(t *testing.T) {
	// 获取无序数组中最大相邻差

	// [2, 6, 3, 4, 5, 10, 9]
	// 9 - 6 = 3
	t.Log(getMaxSortedDistance([]int{2, 6, 3, 4, 5, 10, 9}))
	t.Log(getMaxSortedDistance([]int{1}))

	t.Log(getMaxSortedDistance1([]int{2, 6, 3, 4, 5, 10, 9}))
	t.Log(getMaxSortedDistance1([]int{1}))
}

// 获取无序数组中最大相邻差, 计数排序思想
func getMaxSortedDistance(nums []int) int {
	if len(nums) <= 1 {
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

	// 所有元素相等
	if max == min {
		return 0
	}

	// 计数排序法
	numsCount := make([]int, max-min+1)
	for i := 0; i < len(nums); i++ {
		numsCount[nums[i]-min] = 1
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

// 获取无序数组中最大相邻差, 桶排序思想
func getMaxSortedDistance1(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// 计算最大值和最小值
	max, min := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
			continue
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	// 所有元素相等
	if max == min {
		return 0
	}

	// 初始化桶
	type bucket struct {
		min *int
		max *int
	}
	buckets := make([]bucket, n)

	// 将原数组放入桶
	for i := 0; i < n; i++ {
		index := (nums[i] - min) * (n - 1) / (max - min)
		if buckets[index].min == nil || *buckets[index].min > nums[i] {
			buckets[index].min = &nums[i]
		}
		if buckets[index].max == nil || *buckets[index].max < nums[i] {
			buckets[index].max = &nums[i]
		}
	}

	// 找到最大差值
	var distance int
	leftMax := *buckets[0].max
	for i := 1; i < n; i++ {
		if buckets[i].min == nil {
			continue
		}
		if *buckets[i].min-leftMax > distance {
			distance = *buckets[i].min - leftMax
		}

		leftMax = *buckets[i].max
	}

	return distance
}
