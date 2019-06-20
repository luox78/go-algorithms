package trie_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/trie"
)

func TestTrie(t *testing.T) {
	// 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
	// 示例:
	// Trie trie = new Trie();
	// trie.insert("apple");
	// trie.search("apple");   // 返回 true
	// trie.search("app");     // 返回 false
	// trie.startsWith("app"); // 返回 true
	// trie.insert("app");
	// trie.search("app");     // 返回 true
	// 说明:
	// 你可以假设所有的输入都是由小写字母 a-z 构成的。
	// 保证所有输入均为非空字符串。
	//

	trie := trie.NewTrie()

	trie.Insert("apple")

	// true
	t.Log(trie.Search("apple"))
	// false
	t.Log(trie.Search("app"))
	// true
	t.Log(trie.StartsWith("app"))

	trie.Insert("app")
	// true
	t.Log(trie.Search("app"))
}
