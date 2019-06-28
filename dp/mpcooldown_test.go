package dp_test

import "testing"

func TestMaxProfitWithCoolDown(t *testing.T) {
	// 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

	// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

	// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
	// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

	// 示例:
	// 输入: [1,2,3,0,2]
	// 输出: 3
	// 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
	t.Log(maxProfitWithCoolDown([]int{1, 2, 3, 0, 2}))
}

func maxProfitWithCoolDown(prices []int) int {
	n := len(prices)
	if n <= 0 {
		return 0
	}

	// 状态定义, 1维: 天数, 2维: 0, 未持有; 1, 持有; 2, 冷冻期
	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, 3)
	}
	mp[0][0] = 0
	mp[0][1] = -prices[0]
	mp[0][2] = 0

	var res int
	for i := 1; i < n; i++ {
		// 未持有, 未操作或冷冻期
		if mp[i-1][0] > mp[i-1][2] {
			mp[i][0] = mp[i-1][0]
		} else {
			mp[i][0] = mp[i-1][2]
		}

		// 持有, 未操作或买入
		if mp[i-1][1] > mp[i-1][0]-prices[i] {
			mp[i][1] = mp[i-1][1]
		} else {
			mp[i][1] = mp[i-1][0] - prices[i]
		}

		// 冷冻期, 未操作或卖出
		if mp[i-1][2] > mp[i-1][1]+prices[i] {
			mp[i][2] = mp[i-1][2]
		} else {
			mp[i][2] = mp[i-1][1] + prices[i]
		}

		if mp[i][0] > res {
			res = mp[i][0]
		}
		if mp[i][2] > res {
			res = mp[i][2]
		}
	}
	return res
}
