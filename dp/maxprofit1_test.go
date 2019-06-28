package dp_test

import (
	"testing"
)

func TestMaxProfit1(t *testing.T) {
	// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

	// 如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

	// 注意你不能在买入股票前卖出股票。

	// 示例 1:
	// 输入: [7,1,5,3,6,4]
	// 输出: 5
	// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
	// 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

	t.Log(maxProfit1([]int{7, 1, 5, 3, 6, 4}))

	// 示例 2:
	// 输入: [7,6,4,3,1]
	// 输出: 0
	// 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
	t.Log(maxProfit1([]int{7, 6, 4, 3, 1}))
}

func maxProfit1(prices []int) int {
	n := len(prices)
	if n <= 0 {
		return 0
	}

	// 状态保存, 一维: 天, 二维: 0,未持有; 1,持有
	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, 2)
	}
	mp[0][0] = 0
	mp[0][1] = -prices[0]

	var res int

	for i := 1; i < n; i++ {
		// 持有 0 股, 不动或卖出
		// mp[i][0] = mp[i-1][0]
		// mp[i][0] = mp[i-1][1] + prices[i]
		// fmt.Println(mp[i-1][0], mp[i-1][1]+prices[i])
		if mp[i-1][0] > mp[i-1][1]+prices[i] {
			mp[i][0] = mp[i-1][0]
		} else {
			mp[i][0] = mp[i-1][1] + prices[i]
		}

		// 持有 1 股, 不动或买入
		// mp[i][1] = mp[i-1][1]
		// mp[i][1] = mp[i-1][0] - prices[i]
		if mp[i-1][1] > -prices[i] {
			mp[i][1] = mp[i-1][1]
		} else {
			mp[i][1] = -prices[i]
		}

		if mp[i][0] > res {
			res = mp[i][0]
		}
	}

	return res
}
