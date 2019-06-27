package dp_test

import "testing"

func TestLengthOfLIS(t *testing.T) {
	// 给定一个无序的整数数组，找到其中最长上升子序列的长度。

	// 示例:
	// 输入: [10,9,2,5,3,7,101,18]
	// 输出: 4
	// 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
	// 说明:
	// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
	// 你算法的时间复杂度应该为 O(n^2) 。
	// 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

	// 4
	t.Log(lengthOfLIS(nums))
	// 4
	t.Log(lengthOfLIS1(nums))
}

// 动态规划解法, O(n^2)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	res := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] >= dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}

// 二分解法, O(n log n)
func lengthOfLIS1(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	// 保存最长上升子序列
	res := []int{nums[0]}

	for i := 1; i < n; i++ { // O(n)
		index := getFirstGEIndex(res, nums[i]) // O(log(n))
		if index == len(res) {
			res = append(res, nums[i])
		} else {
			res[index] = nums[i]
		}
	}

	return len(res)
}

// 查找第一个大于等于给定值的元素索引
func getFirstGEIndex(nums []int, num int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1

		if nums[mid] >= num {
			if mid == 0 || nums[mid-1] < num {
				return mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return len(nums)
}
