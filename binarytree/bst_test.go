package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestBinarySearchTree(t *testing.T) {
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

	node := tree.Find(50)
	// 50 34 58
	t.Log(node.Val(), node.Left().Val(), node.Right().Val())

	tree.Delete(50)
}
