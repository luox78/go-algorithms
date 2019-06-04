//
// Package binarytree bst.go
// 二叉搜索树实现
//
package binarytree

// BinarySearchTree 二叉搜索树结构
type BinarySearchTree struct {
	// 根节点
	root *TreeNode
}

// TreeNode 树节点
type TreeNode struct {
	// 值
	val int
	// 左子节点
	left *TreeNode
	// 右子节点
	right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{val: val}
}

// Val 返回节点值
func (n *TreeNode) Val() int { return n.val }

// Left 返回左子节点
func (n *TreeNode) Left() *TreeNode { return n.left }

// Right 返回右子节点
func (n *TreeNode) Right() *TreeNode { return n.right }

// NewBinarySearchTree 新建二叉搜索树
func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

// Find 查找
func (t *BinarySearchTree) Find(val int) *TreeNode {
	p := t.root

	for p != nil {
		if p.val < val {
			p = p.right
		} else if p.val > val {
			p = p.left
		} else {
			return p
		}
	}

	return nil
}

// Insert 插入值
func (t *BinarySearchTree) Insert(val int) {
	if t.root == nil {
		t.root = newTreeNode(val)
		return
	}

	// 获取根结点
	p := t.root

	for p != nil {
		if p.val > val {
			// 在左子树中
			if p.left == nil {
				p.left = newTreeNode(val)
				return
			}
			p = p.left
		} else {
			// 在右子树中
			if p.right == nil {
				p.right = newTreeNode(val)
				return
			}
			p = p.right
		}
	}
}

// Delete 删除值
func (t *BinarySearchTree) Delete(val int) {
	// 要删除的节点
	p := t.root
	// 要删除节点的父节点
	var pp *TreeNode

	for p != nil && p.val != val {
		if p.val > val {
			p = p.left
		} else {
			p = p.right
		}
		pp = p
	}

	// 未找到删除的值
	if p == nil {
		return
	}

	// 要删除的节点下有两个子节点
	if p.left != nil && p.right != nil {
		// 找到右子节点中最小节点
		minP := p.right
		minPP := p
		for minP.left != nil {
			minP = minP.left
			minPP = minP
		}

		// 将最小节点的值替换删除节点的值
		p.val = minP.val

		// 变为删除最小节点
		p = minP
		pp = minPP
	}

	// 删除节点是叶子节点或只有一个子节点
	var child *TreeNode
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	}

	// 删除的为根节点
	if pp == nil {
		t.root = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}

}
