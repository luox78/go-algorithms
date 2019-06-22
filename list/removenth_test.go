package list

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

	// 示例：
	// 给定一个链表: 1->2->3->4->5, 和 n = 2.
	// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
	// 说明：
	// 给定的 n 保证是有效的。

	l := NewSList()
	for i := 1; i < 6; i++ {
		l.PushBack(i)
	}

	// Len: 5, Data: -> 1 -> 2 -> 3 -> 4 -> 5
	t.Log(l)

	l.head = removeNthFromEnd(l.head, 2)
	l.len--

	// Len: 4, Data: -> 1 -> 2 -> 3 -> 5
	t.Log(l)
}

func removeNthFromEnd(head *SListNode, n int) *SListNode {
	if head == nil {
		return nil
	}

	p, q := head, head

	for i := 0; i < n; i++ {
		q = q.next
	}

	if q == nil {
		return p.next
	}

	for q.next != nil {
		q = q.next
		p = p.next
	}

	p.next = p.next.next

	return head
}
