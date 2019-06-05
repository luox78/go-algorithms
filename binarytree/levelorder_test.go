package binarytree_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/queue"

	"github.com/wenjiax/go-algorithms/binarytree"
)

func TestLevelOrder(t *testing.T) {
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

	// 层序遍历

	// [[33] [16 50] [13 18 34 58] [15 17 25 51 66] [19 27 55]]
	t.Log(levelOrder(tree.Find(33)))

	var result [][]int
	// [[33] [16 50] [13 18 34 58] [15 17 25 51 66] [19 27 55]]
	levelOrder1(tree.Find(33), 0, &result)
	t.Log(result)
}

// 层序遍历, 非递归实现
func levelOrder(root *binarytree.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int

	// 利用队列先进先出特性
	q := queue.NewListQueue()
	q.Enqueue(root)

	// 层级
	var level int

	for q.Len() > 0 {
		// 保存当前层级的值
		var levelResult []int
		for i, l := 0, q.Len(); i < l; i++ {
			// 出队列
			p := q.Dequeue().(*binarytree.TreeNode)
			levelResult = append(levelResult, p.Val())

			// 将子节点加入队列
			if p.Left() != nil {
				q.Enqueue(p.Left())
			}
			if p.Right() != nil {
				q.Enqueue(p.Right())
			}
		}

		// 添加进结果
		result = append(result, levelResult)
		level++
	}

	return result
}

// 层序遍历, 递归实现
func levelOrder1(root *binarytree.TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}

	// 初始化
	if len(*result) == level {
		*result = append(*result, []int{})
	}

	(*result)[level] = append((*result)[level], root.Val())

	levelOrder1(root.Left(), level+1, result)
	levelOrder1(root.Right(), level+1, result)
}
