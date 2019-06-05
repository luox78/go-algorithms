package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestPreOrder(t *testing.T) {
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

	// 前序遍历

	// [33 16 13 15 18 17 25 19 27 50 34 58 51 55 66]
	t.Log(preOrder(tree.Find(33)))

	// [33 16 13 15 18 17 25 19 27 50 34 58 51 55 66]
	t.Log(preOrder1(tree.Find(33)))
}

// 前序遍历
func preOrder(root *binarytree.TreeNode) []int {
	if root == nil {
		return nil
	}

	return append(append(append([]int{}, root.Val()), preOrder(root.Left())...), preOrder(root.Right())...)
}

// 前序遍历非递归实现
func preOrder1(root *binarytree.TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int

	// 利用栈先进后出特性
	var stack []*binarytree.TreeNode

	p := root
	for p != nil || len(stack) > 0 {

		// 访问左子节点
		for p != nil {
			result = append(result, p.Val())
			// 加入栈
			stack = append(stack, p)
			p = p.Left()
		}

		// 所有左子节点访问完, 访问节点右子节点
		if len(stack) > 0 {
			// 获取栈元素
			n := len(stack) - 1
			p = stack[n]
			stack = stack[:n]

			p = p.Right()
		}

	}

	return result
}
