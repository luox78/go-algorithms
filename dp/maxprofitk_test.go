package dp_test

import (
	"testing"
)

func TestMaxProfitK(t *testing.T) {
	// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

	// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

	// 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

	// 示例 1:
	// 输入: [2,4,1], k = 2
	// 输出: 2
	// 解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
	t.Log(maxProfitK(2, []int{2, 4, 1}))

	// 示例 2:
	// 输入: [3,2,6,5,0,3], k = 2
	// 输出: 7
	// 解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
	//      随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
	t.Log(maxProfitK(2, []int{3, 2, 6, 5, 0, 3}))
}

// 动态规划求解 k 次股票交易
func maxProfitK(k int, prices []int) int {
	n := len(prices)
	if n <= 0 {
		return 0
	}

	// Leetcode 测试用例 k值过大, 导致用例不通过, 使用贪心算法解决
	if k > n/2 {
		return greedy(prices)
	}

	// 状态定义, 1维: 天数, 2维: 交易次数, 3维: 未持有与持有
	mp := make([][][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([][]int, k+1)
		for kk := 0; kk < k+1; kk++ {
			mp[i][kk] = make([]int, 2)
			if i == 0 {
				mp[i][kk][0] = 0
				mp[i][kk][1] = -prices[0]
			}
		}
	}

	for i := 1; i < n; i++ {
		for kk := 0; kk < k+1; kk++ {
			if kk == 0 {
				// 0次交易状态从前一天过来
				mp[i][kk][0] = mp[i-1][kk][0]
			} else {
				// 持有 0 股, max(前一天未持有, 前一天持有卖出)
				if mp[i-1][kk][0] > mp[i-1][kk-1][1]+prices[i] {
					// 不操作
					mp[i][kk][0] = mp[i-1][kk][0]
				} else {
					// 卖出
					mp[i][kk][0] = mp[i-1][kk-1][1] + prices[i]
				}
			}

			// 持有 1 股, max(前一天持有, 前一天未持有现在买入)
			if mp[i-1][kk][1] > mp[i-1][kk][0]-prices[i] {
				mp[i][kk][1] = mp[i-1][kk][1]
			} else {
				mp[i][kk][1] = mp[i-1][kk][0] - prices[i]
			}
		}
	}
	return mp[n-1][k][0]
}

func greedy(prices []int) int {
	var max int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}
