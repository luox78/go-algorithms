package list

import (
	"testing"
)

func TestMerge2List(t *testing.T) {
	// 合并两个有序链表
	// -> 1 -> 3 -> 5 + -> nil
	// =>
	// -> 1 -> 3 -> 5

	// -> 1 -> 3 -> 5  + -> 2 -> 4
	// =>
	// -> 1 -> 2 -> 3 -> 4 -> 5

	l1 := NewSList()
	l1.PushBack(1)
	l1.PushBack(3)
	l1.PushBack(5)

	// Len: 3, Data: -> 1 -> 3 -> 5
	t.Log(l1)

	l2 := NewSList()
	l2.PushBack(2)
	l2.PushBack(4)

	// Len: 2, Data: -> 2 -> 4
	t.Log(l2)

	l3 := NewSList()
	l3.head = mergeTwoLists(l1.head, l2.head)
	l3.len = l1.Len() + l2.Len()

	// Len: 5, Data: -> 1 -> 2 -> 3 -> 4 -> 5
	t.Log(l3)
}

// 合并两个有序链表
func mergeTwoLists(n1, n2 *SListNode) *SListNode {
	// 退出条件
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	node := new(SListNode)

	// 复制并选择下一个结点
	if n1.Value.(int) < n2.Value.(int) {
		node.Value = n1.Value
		node.next = mergeTwoLists(n1.next, n2)
	} else {
		node.Value = n2.Value
		node.next = mergeTwoLists(n1, n2.next)
	}

	return node
}
