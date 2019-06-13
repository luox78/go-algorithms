package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestMinDepth(t *testing.T) {
	// 	给定一个二叉树，找出其最小深度。

	// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

	// 说明: 叶子节点是指没有子节点的节点。
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
	// 返回它的最小深度 3 。

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

	// 最小深度 3
	t.Log(minDepthDFS(tree.Find(33)))

	// 最小深度 3
	t.Log(minDepthBFS(tree.Find(33)))
}

// 找出二叉树最小深度, DFS 实现
func minDepthDFS(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left() == nil {
		return minDepthDFS(root.Right()) + 1
	}
	if root.Right() == nil {
		return minDepthDFS(root.Left()) + 1
	}

	left, right := minDepthDFS(root.Left()), minDepthDFS(root.Right())
	if left < right {
		return left + 1
	}
	return right + 1
}

// 找出二叉树最小深度, BFS 实现
func minDepthBFS(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*binarytree.TreeNode
	queue = append(queue, root)

	var level int
	for len(queue) > 0 {
		for i, l := 0, len(queue); i < l; i++ {
			p := queue[0]
			queue = queue[1:]

			if p.Left() == nil && p.Right() == nil {
				return level + 1
			}

			if p.Left() != nil {
				queue = append(queue, p.Left())
			}
			if p.Right() != nil {
				queue = append(queue, p.Right())
			}

		}
		level++
	}
	return level
}
