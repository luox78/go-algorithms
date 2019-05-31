package list

import (
	"testing"
)

func TestReversal(t *testing.T) {
	// 单链表反转
	// -> 1 -> 2 -> 3 -> 4 -> 5
	// =>
	// -> 5 -> 4 -> 3 -> 2 -> 1

	l := NewSList()
	for i := 1; i < 6; i++ {
		l.PushBack(i)
	}

	// Len: 5, Data: -> 1 -> 2 -> 3 -> 4 -> 5
	t.Log(l)

	reversal(l)

	// Len: 5, Data: -> 5 -> 4 -> 3 -> 2 -> 1
	t.Log(l)
}

// 单链表反转
func reversal(l *SList) {
	var prev, cur *SListNode
	cur = l.head

	for cur != nil {
		cur.next, cur, prev = prev, cur.next, cur
	}
	l.head = prev
}
