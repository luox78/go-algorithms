package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestMaxDepth(t *testing.T) {
	// 给定一个二叉树，找出其最大深度。

	// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

	// 说明: 叶子节点是指没有子节点的节点。

	// 示例：
	//
	//                         33
	//                 /                \
	//               16                 50
	//          /          \          /      \
	//         13          18        34      58
	//          \        /    \           /     \
	//          15      17    25         51     66
	//                      /    \        \
	//                     19    27       55
	// 返回它的最大深度 5 。

	tree := binarytree.NewBinarySearchTree()
	// 根节点
	tree.Insert(33)
	// 左子树
	tree.Insert(16)
	tree.Insert(13)
	tree.Insert(18)
	tree.Insert(15)
	tree.Insert(17)
	tree.Insert(25)
	tree.Insert(19)
	tree.Insert(27)
	// 右子树
	tree.Insert(50)
	tree.Insert(34)
	tree.Insert(58)
	tree.Insert(51)
	tree.Insert(66)
	tree.Insert(55)

	// 最大深度 5
	t.Log(maxDepthDFS(tree.Find(33)))
}

// 找出二叉树最大深度, DFS 实现
func maxDepthDFS(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := maxDepthDFS(root.Left()), maxDepthDFS(root.Right())
	if left > right {
		return left + 1
	}
	return right + 1
}
