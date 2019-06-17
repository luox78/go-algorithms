package list

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	// 给定一个链表，返回链表开始入环的第一个节点。

	l := NewSList()

	n1 := l.PushBack(1)
	n2 := l.PushBack(2)

	// -> n1(1) -> n2(2) -> n1(1)
	n2.next = n1
	// n1(1)
	t.Log(detectCycle(l.head).Value)

	// -> n1(1) -> n2(2) ->
	n2.next = nil
	// nil
	t.Log(detectCycle(l.head))
}

// 返回入环节点
func detectCycle(head *SListNode) *SListNode {
	if head == nil || head.next == nil {
		return nil
	}

	// 找到第一次重合的节点
	p, q := head, head
	for q != nil && q.next != nil {
		p = p.next
		q = q.next.next

		if p == q {
			break
		}
	}

	// 无环
	if p != q {
		return nil
	}

	// q 指针走的距离为 p 指针的两倍
	// p 再次从头节点, 再次相遇的节点为入环节点
	p = head
	for p != q {
		p = p.next
		q = q.next
	}

	return p
}
