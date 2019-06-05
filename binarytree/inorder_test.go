package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestInOrder(t *testing.T) {
	tree := binarytree.NewBinarySearchTree()
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

	// 中序遍历

	// [13 15 16 17 18 19 25 27 33 34 50 51 55 58 66]
	t.Log(inOrder(tree.Find(33)))

	// [13 15 16 17 18 19 25 27 33 34 50 51 55 58 66]
	t.Log(inOrder1(tree.Find(33)))
}

// 中序遍历, 递归实现
func inOrder(root *binarytree.TreeNode) []int {
	if root == nil {
		return nil
	}

	return append(append(append([]int{}, inOrder(root.Left())...), root.Val()), inOrder(root.Right())...)
}

// 中序遍历, 非递归实现
func inOrder1(root *binarytree.TreeNode) []int {
	var stack []*binarytree.TreeNode
	var result []int

	p := root
	for p != nil || len(stack) > 0 {

		// 访问所有左子节点
		for p != nil {
			// 入栈
			stack = append(stack, p)
			p = p.Left()
		}

		// 左子节点访问完, 访问右子节点
		if len(stack) > 0 {
			// 出栈
			n := len(stack) - 1
			p = stack[n]
			stack = stack[:n]

			result = append(result, p.Val())
			p = p.Right()
		}
	}

	return result
}
