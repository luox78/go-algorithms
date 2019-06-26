package dp_test

import "testing"

func TestMaxProduct(t *testing.T) {
	// 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

	// 示例 1:
	// 输入: [2,3,-2,4]
	// 输出: 6
	// 解释: 子数组 [2,3] 有最大乘积 6。
	t.Log(maxProduct([]int{2, 3, -2, 4}))

	// 示例 2:
	// 输入: [-2,0,-1]
	// 输出: 0
	// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
	t.Log(maxProduct([]int{-2, 0, -1}))
}

func maxProduct(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	// 状态保存, 二维索引 0: 最大值, 1: 最小值(负最大值)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]

	res := dp[0][0]
	for i := 1; i < n; i++ {
		// 状态转移方程
		// dp[i][0] = max(nums[i]*dp[i-1][0], nums[i]*dp[i-1][1], nums[i])
		// dp[i][1] = min(nums[i]*dp[i-1][0], nums[i]*dp[i-1][1], nums[i])

		dp[i][0] = nums[i]
		if nums[i]*dp[i-1][0] > dp[i][0] {
			dp[i][0] = nums[i] * dp[i-1][0]
		}
		if nums[i]*dp[i-1][1] > dp[i][0] {
			dp[i][0] = nums[i] * dp[i-1][1]
		}

		dp[i][1] = nums[i]
		if nums[i]*dp[i-1][0] < dp[i][1] {
			dp[i][1] = nums[i] * dp[i-1][0]
		}
		if nums[i]*dp[i-1][1] < dp[i][1] {
			dp[i][1] = nums[i] * dp[i-1][1]
		}

		if dp[i][0] > res {
			res = dp[i][0]
		}
	}
	return res
}
