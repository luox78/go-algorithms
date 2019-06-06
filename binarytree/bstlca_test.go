package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestLowestCommonAncestor(t *testing.T) {
	// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
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
	t.Log(lowestCommonAncestor(root, p, q).Val())

	// 16
	t.Log(lowestCommonAncestor1(root, p, q).Val())
}

// 二叉搜索树最近公共祖先, 递归实现
func lowestCommonAncestor(root, p, q *binarytree.TreeNode) *binarytree.TreeNode {
	if p.Val() < root.Val() && root.Val() > q.Val() {
		// 在左子树中
		return lowestCommonAncestor(root.Left(), p, q)
	}
	if p.Val() > root.Val() && root.Val() < q.Val() {
		// 在右子树中
		return lowestCommonAncestor(root.Right(), p, q)
	}

	// p, q 在不同子树中
	return root
}

// 二叉搜索树最近公共祖先, 非递归实现
func lowestCommonAncestor1(root, p, q *binarytree.TreeNode) *binarytree.TreeNode {
	r := root
	for r != nil {
		if p.Val() < r.Val() && r.Val() > q.Val() {
			// 在左子树中
			r = r.Left()
		} else if p.Val() > r.Val() && r.Val() < q.Val() {
			// 在右子树中
			r = r.Right()
		} else {
			// 在不同子树中
			return r
		}
	}

	return r
}
