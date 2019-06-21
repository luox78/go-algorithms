package trie_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/trie"
)

func TestFindWords(t *testing.T) {
	// 给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

	// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

	// 示例:
	// 输入:
	// words = ["oath","pea","eat","rain"] and board =
	// [
	//   ['o','a','a','n'],
	//   ['e','t','a','e'],
	//   ['i','h','k','r'],
	//   ['i','f','l','v']
	// ]

	// 输出: ["eat","oath"]
	// 说明:
	// 你可以假设所有输入都由小写字母 a-z 组成。

	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}
	// [oath eat]
	t.Log(findWords(board, []string{"oath", "pea", "eat", "rain"}))

	board = [][]byte{
		[]byte{'a', 'b'},
		[]byte{'a', 'a'},
	}
	// [aaa aaab aba baa aaba]
	t.Log(findWords(board, []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}))
}

func findWords(board [][]byte, words []string) []string {
	if len(board) <= 0 || len(words) <= 0 {
		return nil
	}

	// 将单词加入 trie
	trie := trie.NewTrie()

	for _, w := range words {
		trie.Insert(w)
	}

	m, n := len(board), len(board[0])

	// 记录被访问
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var res []string

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			find(board, "", i, j, visited, trie, &res)
		}
	}

	return res
}

func find(board [][]byte, curWord string, x, y int, visited [][]bool, trie trie.Trie, res *[]string) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || visited[x][y] {
		return
	}

	curWord += string(board[x][y])

	if !trie.StartsWith(curWord) {
		return
	}

	if trie.Search(curWord) && !exist(*res, curWord) {
		*res = append(*res, curWord)
	}

	visited[x][y] = true

	find(board, curWord, x-1, y, visited, trie, res)
	find(board, curWord, x+1, y, visited, trie, res)
	find(board, curWord, x, y-1, visited, trie, res)
	find(board, curWord, x, y+1, visited, trie, res)

	visited[x][y] = false
}

func exist(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}
