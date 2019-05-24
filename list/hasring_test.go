package list

import (
	"testing"
)

func TestHasRing(t *testing.T) {
	// 判断链表中是否有环
	// -> n1 -> n2 -> n1
	// => true

	l := NewSList()

	n1 := l.PushBack(1)
	n2 := l.PushBack(2)

	// -> n1 -> n2 -> n1
	n2.next = n1
	// true
	t.Log(hasRing(l))

	// // -> n1 -> n2 ->
	n2.next = nil
	// false
	t.Log(hasRing(l))
}

// 判断链表中是否有环
func hasRing(l *SList) bool {
	p, q := l.Front(), l.Front()

	for q != nil && q.Next() != nil {
		p = p.Next()
		q = q.Next().Next()

		if p == q {
			return true
		}
	}

	return false
}
