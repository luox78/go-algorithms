package list_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/list"
)

func TestGetMidNode(t *testing.T) {
	// 求链表的中间结点
	// -> 1 -> 2 -> 3
	// =>
	// 2
	// -> 1 -> 2 -> 3 -> 4
	// =>
	// 3
	l := list.NewSList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// 2
	t.Log(getMidNode(l).Value)

	// 3
	l.PushBack(4)
	t.Log(getMidNode(l).Value)

	// 3
	l.PushBack(5)
	t.Log(getMidNode(l).Value)
}

// 求链表的中间结点
func getMidNode(l *list.SList) *list.SListNode {
	p, q := l.Front(), l.Front()

	for q != nil && q.Next() != nil {
		p = p.Next()
		q = q.Next().Next()
	}
	return p
}
