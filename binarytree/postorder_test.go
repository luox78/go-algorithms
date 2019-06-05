package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestPostOrder(t *testing.T) {
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

	// 后序遍历

	// [15 13 17 19 27 25 18 16 34 55 51 66 58 50 33]
	t.Log(postOrder(tree.Find(33)))

	// [15 13 17 19 27 25 18 16 34 55 51 66 58 50 33]
	t.Log(postOrder1(tree.Find(33)))
}

// 后序遍历, 递归实现
func postOrder(root *binarytree.TreeNode) []int {
	if root == nil {
		return nil
	}

	return append(append(append([]int{}, postOrder(root.Left())...), postOrder(root.Right())...), root.Val())
}

// 后序遍历, 非递归实现
func postOrder1(root *binarytree.TreeNode) []int {
	var result []int
	var stack []*binarytree.TreeNode

	// 记录最后访问的节点
	var last *binarytree.TreeNode

	p := root

	for p != nil || len(stack) > 0 {
		// 所有左子节点入栈
		for p != nil {
			stack = append(stack, p)
			p = p.Left()
		}

		if len(stack) > 0 {
			p = stack[len(stack)-1]

			// 判断是否为叶子节点或访问过的节点
			if p.Right() == nil || p.Right() == last {
				result = append(result, p.Val())
				// 出栈
				n := len(stack) - 1
				stack = stack[:n]
				last = p
				p = nil
			} else {
				p = p.Right()
			}
		}
	}

	return result
}
