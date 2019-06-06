package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestBinaryTreeLCA(t *testing.T) {
	// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
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

	root := tree.Find(33)
	p := tree.Find(15)
	q := tree.Find(19)

	// 16
	t.Log(binaryTreeLCA(root, p, q).Val())
}

// 二叉树最近公共祖先
func binaryTreeLCA(root, p, q *binarytree.TreeNode) *binarytree.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// 在左右子树中查找
	left := binaryTreeLCA(root.Left(), p, q)
	right := binaryTreeLCA(root.Right(), p, q)

	// 在右子树
	if left == nil {
		return right
	}
	// 在左子树
	if right == nil {
		return left
	}

	// 在不同子树
	return root
}
