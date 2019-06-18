package list_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/list"
)

func TestLastKNode(t *testing.T) {
	// 获取链表中倒数第 K 个节点
	l := list.NewSList()

	for i := 1; i < 10; i++ {
		l.PushBack(i)
	}
	t.Log(l)

	// 7
	t.Log(getLastKNode(l, 3).Value)

	// 9
	t.Log(getLastKNode(l, 1).Value)

	// 1
	t.Log(getLastKNode(l, 9).Value)
}

// 获取链表中倒数第 K 个节点
func getLastKNode(l *list.SList, k int) *list.SListNode {
	if l == nil || l.Len() < k {
		return nil
	}

	// 双指针
	p, q := l.Front(), l.Front()

	// 快指针先走 k 步
	for i := 0; i < k; i++ {
		q = q.Next()
	}

	for q != nil {
		p = p.Next()
		q = q.Next()
	}
	return p
}
