package bit_test

import "testing"

func TestTotalNQueens(t *testing.T) {
	// 给定一个整数 n，返回 n 皇后不同的解决方案的数量。

	// 示例:
	// 输入: 4
	// 输出: 2
	// 解释: 4 皇后问题存在如下两个不同的解法。
	// [
	//  [".Q..",  // 解法 1
	//   "...Q",
	//   "Q...",
	//   "..Q."],

	//  ["..Q.",  // 解法 2
	//   "Q...",
	//   "...Q",
	//   ".Q.."]
	// ]

	t.Log(totalNQueens(4))
	t.Log(totalNQueens(8))
	t.Log(totalNQueens(1))
}

func totalNQueens(n int) int {
	var count int
	bitDFS(n, 0, 0, 0, 0, &count)
	return count
}

func bitDFS(n, row, col, left, right int, count *int) {
	// 递归终止条件
	if row >= n {
		*count++
		return
	}

	// 获取所有可以放的位
	bits := (^(col | left | right)) & ((1 << uint(n)) - 1)

	for bits > 0 {
		// 得到1个空位
		p := bits & -bits
		bitDFS(n, row+1, col|p, (left|p)<<1, (right|p)>>1, count)
		// 去掉放置的空位
		bits &= bits - 1
	}
}
