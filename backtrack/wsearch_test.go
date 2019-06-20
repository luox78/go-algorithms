package backtrack_test

import (
	"testing"
)

func TestWordExist(t *testing.T) {
	// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。

	// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

	// 示例:
	// board =
	// [
	//   ['A','B','C','E'],
	//   ['S','F','C','S'],
	//   ['A','D','E','E']
	// ]
	// 给定 word = "ABCCED", 返回 true.
	// 给定 word = "SEE", 返回 true.
	// 给定 word = "ABCB", 返回 false.

	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}

	// true
	t.Log(wordExist(board, "ABCCED"))
	// true
	t.Log(wordExist(board, "SEE"))
	// false
	t.Log(wordExist(board, "ABCB"))
}

// 返回给定单词是否存在
func wordExist(board [][]byte, word string) bool {
	if len(board) <= 0 {
		return false
	}

	m, n := len(board), len(board[0])

	// 记录被访问
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 单词第一个前缀匹配则继续判断
			if board[i][j] == word[0] && backtrack2(board, word, 0, i, j, visited) {
				return true
			}
		}
	}

	return false
}

func backtrack2(board [][]byte, word string, index, i, j int, visited [][]bool) bool {
	if index == len(word) {
		return true
	}

	// 剪枝
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		board[i][j] != word[index] || visited[i][j] {
		return false
	}

	visited[i][j] = true
	// 判断相邻位置
	if backtrack2(board, word, index+1, i-1, j, visited) ||
		backtrack2(board, word, index+1, i+1, j, visited) ||
		backtrack2(board, word, index+1, i, j-1, visited) ||
		backtrack2(board, word, index+1, i, j+1, visited) {
		return true
	}
	// 回溯
	visited[i][j] = false

	return false
}
