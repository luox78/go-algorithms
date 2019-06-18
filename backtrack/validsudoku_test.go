package backtrack_test

import "testing"

func TestValidSuduku(t *testing.T) {
	// 判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

	// 数字 1-9 在每一行只能出现一次。
	// 数字 1-9 在每一列只能出现一次。
	// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
	// 数独部分空格内已填入了数字，空白格用 '.' 表示。

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

	// true
	t.Log(validSudoku(board))

	// false
	board[8][8] = '1'
	t.Log(validSudoku(board))
}

// 验证已经填入的数字是否有效
func validSudoku(board [][]byte) bool {
	// 行
	for row := 0; row < 9; row++ {
		// 列
		for col := 0; col < 9; col++ {
			// 未填充值
			if board[row][col] == '.' {
				continue
			}

			if !isValid1(board, row, col) {
				return false
			}
		}
	}

	return true
}

func isValid1(board [][]byte, row, col int) bool {
	num := board[row][col]

	for i := 0; i < 9; i++ {
		// 判断每行当前列中是否存在
		if i != row && board[i][col] == num {
			return false
		}
		// 判断当前行中每列是否存在
		if i != col && board[row][i] == num {
			return false
		}
		// 判断 3*3 区域中是否存在
		if (3*(row/3)+i/3) != row && (3*(col/3)+i%3) != col && board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}
