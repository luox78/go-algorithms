package binarytree_test

import (
	"math"
	"testing"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestIsValidBST(t *testing.T) {
	// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。

	// 假设一个二叉搜索树具有如下特征：

	// 节点的左子树只包含小于当前节点的数。
	// 节点的右子树只包含大于当前节点的数。
	// 所有左子树和右子树自身必须也是二叉搜索树。

	// 示例 1:
	// 输入:
	//     2
	//    / \
	//   1   3
	// 输出: true
	tree := binarytree.NewBinarySearchTree()
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)
	t.Log(isValidBST(tree.Find(2), math.MinInt64, math.MaxInt64))

	// true
	tree = binarytree.NewBinarySearchTree()
	tree.Insert(2147483647)
	t.Log(isValidBST(tree.Find(2147483647), math.MinInt64, math.MaxInt64))

}

func isValidBST(root *binarytree.TreeNode, min, max int) bool {
	// 空树为二叉搜索树
	if root == nil {
		return true
	}

	// 判断左子节点
	if root.Val() <= min {
		return false
	}

	// 判断右子节点
	if root.Val() >= max {
		return false
	}

	return isValidBST(root.Left(), min, root.Val()) &&
		isValidBST(root.Right(), root.Val(), max)
}
