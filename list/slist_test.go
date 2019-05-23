package list_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/list"
)

func TestSList(t *testing.T) {
	l := list.NewSList()

	// Len: 0, Data:
	t.Log(l)

	// Len: 1, Data: -> 1
	l.PushBack(1)
	t.Log(l)

	// Len: 2, Data: -> 2 -> 1
	v := l.PushFront(2)
	t.Log(l)

	// Len: 1, Data: -> 1
	l.Remove(v)
	t.Log(l)

	// Len: 2, Data: -> 1 -> 2
	v = l.PushBack(2)
	t.Log(l)

	// Len: 3, Data: -> 1 -> 2 -> 3
	l.InsertAfter(3, v)
	t.Log(l)

	// Len: 4, Data: -> 1 -> 2 -> 4 -> 3
	l.InsertBefore(4, v)
	t.Log(l)

	// 遍历
	for p := l.Front(); p != nil; p = p.Next() {
		t.Log(p.Value)
	}
}
