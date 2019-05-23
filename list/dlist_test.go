package list_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/list"
)

func TestDList(t *testing.T) {
	l := list.NewDList()

	// Len: 0, Data:
	t.Log(l)

	// Len: 1, Data: <-> 1
	v := l.PushBack(1)
	t.Log(l)

	// Len: 2, Data: <-> 2 <-> 1
	l.PushFront(2)
	t.Log(l)

	// Len: 3, Data: <-> 2 <-> 1 <-> 3
	l.PushBack(3)
	t.Log(l)

	// Len: 2, Data: <-> 2 <-> 3
	l.Remove(v)
	t.Log(l)

	// Len: 3, Data: <-> 2 <-> 3 <-> 5
	v = l.PushBack(5)
	t.Log(l)

	// Len: 4, Data: <-> 2 <-> 3 <-> 5 <-> 6
	l.InsertAfter(6, v)
	t.Log(l)

	// Len: 5, Data: <-> 2 <-> 3 <-> 4 <-> 5 <-> 6
	l.InsertBefore(4, v)
	t.Log(l)

	// 从头结点向后遍历
	for p := l.Front(); p != nil; p = p.Next() {
		t.Log(p.Value)
	}

	// 从尾结点向前遍历
	for p := l.Back(); p != nil; p = p.Prev() {
		t.Log(p.Value)
	}
}
