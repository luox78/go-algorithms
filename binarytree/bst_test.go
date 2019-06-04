package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestBinarySearchTree(t *testing.T) {
	tree := binarytree.NewBinarySearchTree()
	//
	//							33
	// 				   /				\
	// 			     16					50
	// 			/		  \			  /		\
	// 		   13		  18		34		 58
	// 			\		/	 \			   /	\
	// 			15	   17	  25		 51		66
	// 						/	\			\
	//					   19	27			55

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

	var result []int
	preOrder(tree.Find(33), &result)
	// [33 16 13 15 18 17 25 19 27 50 34 58 51 55 66]
	t.Log(result)
}

// 前序遍历
func preOrder(p *binarytree.TreeNode, result *[]int) {
	if p == nil {
		return
	}

	*result = append(*result, p.Val())

	preOrder(p.Left(), result)
	preOrder(p.Right(), result)
}
