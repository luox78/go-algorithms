package list

import (
	"testing"
)

func TestRingLen(t *testing.T) {
	// 获取链表中环的长度

	// -> 1 -> 2(n1) -> 3 -> 4(n2) -> 2(n1)
	l := NewSList()

	l.PushBack(1)
	n1 := l.PushBack(2)
	l.PushBack(3)
	n2 := l.PushBack(4)

	n2.next = n1

	// 3
	t.Log(getRingLen(l))
}

// 获取链表中环的长度
func getRingLen(l *SList) int {
	if l == nil {
		return 0
	}

	p, q := l.Front(), l.Front()

	// 找到环的位置
	for q != nil && q.Next() != nil {
		p = p.Next()
		q = q.Next().Next()

		if p == q {
			break
		}
	}

	// 没有环
	if p != q {
		return 0
	}

	// 计算环的长度
	p = p.Next()
	len := 1

	for p != q {
		p = p.Next()
		len++
	}
	return len
}
