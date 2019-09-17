package backtrack_test

import "testing"

func TestLongestIncreasingPath(t *testing.T) {
	// 给定一个整数矩阵，找出最长递增路径的长度。

	// 对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。

	// 示例 1:
	// 输入: nums =
	// [
	//   [9,9,4],
	//   [6,6,8],
	//   [2,1,1]
	// ]
	// 输出: 4
	// 解释: 最长递增路径为 [1, 2, 6, 9]。
	t.Log(longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}))

	// 示例 2:
	// 输入: nums =
	// [
	//   [3,4,5],
	//   [3,2,6],
	//   [2,2,1]
	// ]
	// 输出: 4
	// 解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
	t.Log(longestIncreasingPath([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}))
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	cache := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		cache[i] = make([]int, len(matrix[0]))
	}

	var max int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			res := getMax(matrix, row, col, cache)
			if res > max {
				max = res
			}
		}
	}

	return max
}

var xy = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// dsf + 记忆法
func getMax(matrix [][]int, row, col int, cache [][]int) int {
	if cache[row][col] != 0 {
		return cache[row][col]
	}

	for i := 0; i < len(xy); i++ {
		x, y := row+xy[i][0], col+xy[i][1]

		if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[row][col] < matrix[x][y] {

			if m := getMax(matrix, x, y, cache); m > cache[row][col] {
				cache[row][col] = m

			}
		}
	}
	cache[row][col]++
	return cache[row][col]
}
