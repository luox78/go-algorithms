package binarytree

import (
	"testing"
)

func TestInvertTree(t *testing.T) {
	// 翻转一棵二叉树。

	// 示例：

	// 输入：
	//      4
	//    /   \
	//   2     7
	//  / \   / \
	// 1   3 6   9

	// 输出：
	//      4
	//    /   \
	//   7     2
	//  / \   / \
	// 9   6 3   1

	tree := NewBinarySearchTree()
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(6)
	tree.Insert(9)

	// 翻转前：[1 2 3 4 6 7 9]
	t.Log(inOrder(tree.root))

	// 翻转
	root := invertTree(tree.root)

	// 翻转后：[9 7 6 4 3 2 1]
	t.Log(inOrder(root))
}

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.left, root.right = invertTree(root.right), invertTree(root.left)

	return root
}

// 中序遍历
func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	return append(append(append([]int{}, inOrder(root.left)...), root.val), inOrder(root.right)...)
}
