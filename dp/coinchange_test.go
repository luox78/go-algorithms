package dp_test

import "testing"

func TestCoinChange(t *testing.T) {
	// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

	// 示例 1:
	// 输入: coins = [1, 2, 5], amount = 11
	// 输出: 3
	// 解释: 11 = 5 + 5 + 1
	t.Log(coinChange([]int{1, 2, 5}, 11))

	// 示例 2:
	// 输入: coins = [2], amount = 3
	// 输出: -1
	t.Log(coinChange([]int{2}, 3))

	// 说明:
	// 你可以认为每种硬币的数量是无限的。
}

// 动态规划求解, O(n*m)
func coinChange(coins []int, amount int) int {
	// 状态定义
	count := make([]int, amount+1)

	// i : 总金额为 i 时所需的硬币个数
	for i := 1; i <= amount; i++ {
		// 默认值, 返回时判断是否可以凑成所需金额
		count[i] = amount + 1

		// 状态转移方程, count[i] = min(count[i - coins[0...j]]+1)
		// 当前硬币个数等于减去目前硬币面额之后的硬币个数加一
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				if count[i] > count[i-coins[j]]+1 {
					count[i] = count[i-coins[j]] + 1
				}
			}
		}
	}

	// 无法凑成所需金额
	if count[amount] > amount {
		return -1
	}
	return count[amount]
}
