package backtrack_test

import (
	"fmt"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	// 编写一个程序，通过已填充的空格来解决数独问题。

	// 一个数独的解法需遵循如下规则：
	// 数字 1-9 在每一行只能出现一次。
	// 数字 1-9 在每一列只能出现一次。
	// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
	// 空白格用 '.' 表示。
	// Note:
	// 给定的数独序列只包含数字 1-9 和字符 '.' 。
	// 你可以假设给定的数独只有唯一解。
	// 给定数独永远是 9x9 形式的。

	board := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	solveSudoku(board)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%s\t", string(board[i][j]))
		}
		fmt.Println()
	}
}

// 求解数独
func solveSudoku(board [][]byte) {
	if len(board) <= 0 {
		return
	}
	solve(board)
}

// 回溯算法求解
func solve(board [][]byte) bool {
	// 行
	for row := 0; row < 9; row++ {
		// 列
		for col := 0; col < 9; col++ {
			// 固定值
			if board[row][col] != '.' {
				continue
			}

			for num := byte('1'); num <= byte('9'); num++ {
				if isValid(board, row, col, num) {
					board[row][col] = byte(num)

					// 继续下一层
					if solve(board) {
						return true
					} else {
						// 回溯到当前状态
						board[row][col] = '.'
					}
				}
			}
			// 所有数字都不合法
			return false
		}

	}

	return true
}

// 判断填充的数字是否有效
func isValid(board [][]byte, row, col int, num byte) bool {

	for i := 0; i < 9; i++ {
		// 判断每行当前列中是否存在
		if board[i][col] == num {
			return false
		}
		// 判断当前行中每列是否存在
		if board[row][i] == num {
			return false
		}

		// 判断每个 3*3 区域是否存在
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}
