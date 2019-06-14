package binarytree

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	// 给定一个二叉树，原地将它展开为链表。

	// 例如，给定二叉树
	//     4
	//    / \
	//   2   5
	//  / \   \
	// 1   3   6

	// 将其展开为：
	//  4
	//  \
	//   2
	//    \
	//    1
	//     \
	//     3
	//      \
	//       5
	//       \
	//        6
	tree := NewBinarySearchTree()
	tree.Insert(4)

	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)

	tree.Insert(5)
	tree.Insert(6)

	// [4 2 1 3 5 6]
	t.Log(preOrder(tree.Find(4)))
	// [1 2 3 4 5 6]
	t.Log(inOrder(tree.Find(4)))

	flatten(tree.Find(4))

	// [4 2 1 3 5 6]
	t.Log(preOrder(tree.Find(4)))
	// [4 2 1 3 5 6]
	t.Log(inOrder(tree.Find(4)))
}

// 将二叉树展开为链表, 后序遍历
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.left)
	flatten(root.right)

	tmp := root.right
	root.right = root.left
	root.left = nil

	p := root
	for p.right != nil {
		p = p.right
	}
	p.right = tmp
}

func preOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(append(append([]int{}, root.val), preOrder(root.left)...), preOrder(root.right)...)
}
