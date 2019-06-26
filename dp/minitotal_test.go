package dp_test

import "testing"

func TestMinimumTotal(t *testing.T) {
	// 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

	// 例如，给定三角形：
	// [
	//      [2],
	//     [3,4],
	//    [6,5,7],
	//   [4,1,8,3]
	// ]
	// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

	// 说明：
	// 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	t.Log(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n <= 0 {
		return 0
	}

	// 保存前一层状态
	state := triangle[n-1]

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			min := state[j]
			if min > state[j+1] {
				min = state[j+1]
			}
			// state[i][j] = min(state[i+1,j], state[i+1,j+1]) + triangle[i][j]
			state[j] = triangle[i][j] + min
		}
	}
	return state[0]
}
