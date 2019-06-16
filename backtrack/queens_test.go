package backtrack_test

import (
	"testing"
)

func TestNQueens(t *testing.T) {
	// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

	// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

	// 示例:
	// 输入: 4
	// 输出: [
	//  [".Q..",  // 解法 1
	//   "...Q",
	//   "Q...",
	//   "..Q."],

	//  ["..Q.",  // 解法 2
	//   "Q...",
	//   "...Q",
	//   ".Q.."]
	// ]
	// 解释: 4 皇后问题存在两个不同的解法。
	solveNQueens(7)
	// res := solveNQueens(7)
	// for _, r := range res {
	// 	t.Log(r)
	// }
}

// 求解 n 皇后问题
func solveNQueens(n int) [][]string {
	var result [][]int

	backtrack(0, n, []int{}, &result)

	return generateResult(result, n)
}

// 回溯算法求解
func backtrack(row, n int, queens []int, result *[][]int) {
	if row >= n {
		// 找到一种解法
		// 避免添加的值被覆盖
		tmp := make([]int, len(queens))
		copy(tmp, queens)

		*result = append(*result, tmp)

		return
	}

	for col := 0; col < n; col++ {
		if isOK(row, col, n, queens) {
			backtrack(row+1, n, append(queens, col), result)
		}
	}
}

// 判断当前位置是否符合条件
func isOK(row, col, n int, queens []int) bool {
	// 对角线位置
	leftUp, rightUp := col-1, col+1
	// 从下往上逐行判断
	for i := row - 1; i >= 0; i-- {

		if queens[i] == col {
			return false
		}
		if leftUp >= 0 && queens[i] == leftUp {
			return false
		}
		if rightUp < n && queens[i] == rightUp {
			return false
		}

		leftUp--
		rightUp++
	}
	return true
}

func generateResult(queens [][]int, n int) [][]string {
	var result [][]string

	for _, res := range queens {
		var r []string
		for _, q := range res {
			var s string
			for i := 0; i < n; i++ {
				if q == i {
					s += "Q"
				} else {
					s += "."
				}
			}
			r = append(r, s)
		}
		result = append(result, r)
	}

	return result
}
