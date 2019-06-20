// Package trie trie.go
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
package trie

// Node Trie节点结构
type Node struct {
	// 子节点
	children []*Node
	// 当前节点值
	val rune
	// 是否构成单词
	isWord bool
}

// NewTrieNode 返回新建 Trie节点
func NewTrieNode(val rune) *Node {
	return &Node{
		children: make([]*Node, 26), // 26 个小写字母构成
		val:      val,
	}
}

// Trie 前缀树结构
type Trie struct {
	root *Node
}

// NewTrie 新建 Trie
func NewTrie() Trie {
	return Trie{
		root: NewTrieNode(' '),
	}
}

// Insert 插入单词
func (t *Trie) Insert(word string) {
	if t.root == nil {
		trie := NewTrie()
		t = &trie
	}
	ws := t.root

	for _, c := range word {
		// 没有找到前缀
		c = c - 'a'
		if ws.children[c] == nil {
			ws.children[c] = NewTrieNode(c)
		}

		ws = ws.children[c]
	}
	ws.isWord = true
}

// Search 返回是否存在给定的单词
func (t *Trie) Search(word string) bool {
	if t.root == nil {
		return false
	}
	ws := t.root

	for _, c := range word {
		c = c - 'a'
		if ws.children[c] == nil {
			return false
		}

		ws = ws.children[c]
	}

	return ws.isWord
}

// StartsWith 返回是否存在给定的前缀
func (t *Trie) StartsWith(prefix string) bool {
	if t.root == nil {
		return false
	}
	ws := t.root

	for _, c := range prefix {
		c = c - 'a'
		if ws.children[c] == nil {
			return false
		}

		ws = ws.children[c]
	}
	return true
}
