package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/queue"

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

	// 前序遍历
	var result1 []int
	preOrder(tree.Find(33), &result1)
	// [33 16 13 15 18 17 25 19 27 50 34 58 51 55 66]
	t.Log(result1)

	// 中序遍历
	var result2 []int
	inOrder(tree.Find(33), &result2)
	// [13 15 16 17 18 19 25 27 33 34 50 51 55 58 66]
	t.Log(result2)

	// 后序遍历
	var result3 []int
	postOrder(tree.Find(33), &result3)
	// [15 13 17 19 27 25 18 16 34 55 51 66 58 50 33]
	t.Log(result3)

	// 层序遍历
	var result4 []int
	levelOrder(tree.Find(33), &result4)
	// [33 16 50 13 18 34 58 15 17 25 51 66 19 27 55]
	t.Log(result4)
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

// 中序遍历
func inOrder(p *binarytree.TreeNode, result *[]int) {
	if p == nil {
		return
	}

	inOrder(p.Left(), result)

	*result = append(*result, p.Val())

	inOrder(p.Right(), result)
}

// 后序遍历
func postOrder(p *binarytree.TreeNode, result *[]int) {
	if p == nil {
		return
	}

	postOrder(p.Left(), result)
	postOrder(p.Right(), result)

	*result = append(*result, p.Val())
}

// 层序遍历
func levelOrder(p *binarytree.TreeNode, result *[]int) {
	if p == nil {
		return
	}

	// 借助队列先进先出特性
	q := queue.NewListQueue()
	q.Enqueue(p)

	for q.Len() > 0 {
		// 从队列取出节点
		node := q.Dequeue().(*binarytree.TreeNode)

		*result = append(*result, node.Val())

		// 将子节点入队列
		if n := node.Left(); n != nil {
			q.Enqueue(n)
		}
		if n := node.Right(); n != nil {
			q.Enqueue(n)
		}
	}

}
