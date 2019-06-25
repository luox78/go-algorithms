package binarytree

import (
	"math"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	// 给定一个非空二叉树，返回其最大路径和。

	// 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

	// 示例 1:
	// 输入:
	//        1
	//       / \
	//      2   3
	// 输出: 6
	tree := NewBinarySearchTree()
	tree.root = newTreeNode(1)
	tree.root.left = newTreeNode(2)
	tree.root.right = newTreeNode(3)

	max := math.MinInt32
	maxPathSum(tree.root, &max)
	t.Log(max)

	// 示例 2:
	// 输入:
	//    -10
	//    / \
	//   9  20
	//     /  \
	//    15   7
	// 输出: 42
	tree = NewBinarySearchTree()
	tree.root = newTreeNode(-10)
	tree.root.left = newTreeNode(9)
	tree.root.right = newTreeNode(20)
	tree.root.right.left = newTreeNode(15)
	tree.root.right.right = newTreeNode(7)

	max = math.MinInt32
	maxPathSum(tree.root, &max)
	t.Log(max)
}

func maxPathSum(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := maxPathSum(root.left, max)
	if left < 0 {
		left = 0
	}
	right := maxPathSum(root.right, max)
	if right < 0 {
		right = 0
	}

	if root.val+left+right > *max {
		*max = root.val + left + right
	}

	if left > right {
		return root.val + left
	}
	return root.val + right
}
