package backtrack_test

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
	t.Log(totalNQueens(1))
}

// 返回 n 皇后不同的解决方案的数量
func totalNQueens(n int) int {
	var total int

	backtrack1(0, n, &total)
	return total
}

// 保存列, 左对角线, 右对角线是否有皇后状态标记位
var cols, leftup, rightup = make(map[int]bool), make(map[int]bool), make(map[int]bool)

// 回溯算法求解
func backtrack1(row, n int, total *int) {
	if row >= n {
		*total++
		return
	}

	for col := 0; col < n; col++ {
		if cols[col] || leftup[row+col] || rightup[row-col] {
			continue
		}

		// 保存状态
		cols[col] = true
		leftup[row+col] = true
		rightup[row-col] = true

		backtrack1(row+1, n, total)

		// 清除状态
		delete(cols, col)
		delete(leftup, row+col)
		delete(rightup, row-col)
	}
}
