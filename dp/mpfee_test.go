package dp_test

import "testing"

func TestMaxProfitWithFee(t *testing.T) {
	// 给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。

	// 你可以无限次地完成交易，但是你每次交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

	// 返回获得利润的最大值。

	// 示例 1:
	// 输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
	// 输出: 8
	// 解释: 能够达到的最大利润:
	// 在此处买入 prices[0] = 1
	// 在此处卖出 prices[3] = 8
	// 在此处买入 prices[4] = 4
	// 在此处卖出 prices[5] = 9
	// 总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
	t.Log(maxProfitWithFee([]int{1, 3, 2, 8, 4, 9}, 2))
}

func maxProfitWithFee(prices []int, fee int) int {
	n := len(prices)
	if n <= 0 {
		return 0
	}

	// 状态定义, 1维: 天数, 2维: 是否持有
	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, 2)
	}
	mp[0][0] = 0
	mp[0][1] = -prices[0] - fee

	var res int
	for i := 1; i < n; i++ {
		// 未持有, 未操作或卖出
		if mp[i-1][0] > mp[i-1][1]+prices[i] {
			mp[i][0] = mp[i-1][0]
		} else {
			mp[i][0] = mp[i-1][1] + prices[i]
		}

		// 持有, 未操作或买入
		if mp[i-1][1] > mp[i-1][0]-prices[i]-fee {
			mp[i][1] = mp[i-1][1]
		} else {
			mp[i][1] = mp[i-1][0] - prices[i] - fee
		}

		if mp[i][0] > res {
			res = mp[i][0]
		}
	}
	return res
}
