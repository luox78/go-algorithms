package list

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	// 合并 k 个排序链表，返回合并后的排序链表。

	// 示例:

	// 输入:
	// [
	//   1->4->5,
	//   1->3->4,
	//   2->6
	// ]
	// 输出: 1->1->2->3->4->4->5->6

	l1 := NewSList()
	l1.PushBack(1)
	l1.PushBack(4)
	l1.PushBack(5)

	l2 := NewSList()
	l2.PushBack(1)
	l2.PushBack(3)
	l2.PushBack(4)

	l3 := NewSList()
	l3.PushBack(2)
	l3.PushBack(6)

	n := mergeKLists([]*SListNode{l1.head, l2.head, l3.head})
	l := NewSList()
	l.head = n
	l.len = l1.Len() + l2.Len() + l3.Len()
	t.Log(l)

}

// mergeKLists 合并 k 个排序链表, 分治法
func mergeKLists(lists []*SListNode) *SListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*SListNode, left, right int) *SListNode {
	if left == right {
		return lists[left]
	}

	mid := left + (right-left)/2
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)

	return mergeTwo(l1, l2)
}

// 合并两个链表
func mergeTwo(l1, l2 *SListNode) *SListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Value.(int) < l2.Value.(int) {
		l1.next = mergeTwo(l1.next, l2)
		return l1
	}

	l2.next = mergeTwo(l1, l2.next)
	return l2
}
