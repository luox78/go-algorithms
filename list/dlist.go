//
// Package list dlist.go
// 双向链表实现
//
package list

import "fmt"

// DListNode 双向链表结点结构
type DListNode struct {
	// 数据
	Value interface{}
	// 前驱结点
	prev *DListNode
	// 后继结点
	next *DListNode
}

// Prev 获取前驱结点
func (n *DListNode) Prev() *DListNode { return n.prev }

// Next 获取后继结点
func (n *DListNode) Next() *DListNode { return n.next }

// DList 双向链表
type DList struct {
	head *DListNode
	len  int
}

// NewDList 返回双向链表
func NewDList() *DList {
	return new(DList)
}

// Len 返回双向链表长度
func (l *DList) Len() int { return l.len }

// Front 返回头部结点
func (l *DList) Front() *DListNode {
	return l.head
}

// Back 返回尾部结点
func (l *DList) Back() *DListNode {
	if l.len == 0 {
		return nil
	}

	p := l.head
	for p.next != nil {
		p = p.next
	}

	return p
}

// PushFront 添加至头结点
func (l *DList) PushFront(v interface{}) *DListNode {
	n := &DListNode{Value: v}

	if l.len > 0 {
		l.head.prev = n
		n.next = l.head
	}

	l.head = n
	l.len++
	return l.head
}

// PushBack 添加至尾结点
func (l *DList) PushBack(v interface{}) *DListNode {
	if l.len == 0 {
		l.head = &DListNode{Value: v}
		l.len++
		return l.head
	}

	// 查找尾结点
	p := l.head
	for p.next != nil {
		p = p.next
	}

	return l.insertVal(v, p)
}

// 插入结点到指定结点后
func (l *DList) insert(v, at *DListNode) *DListNode {
	p := at.next
	at.next = v

	v.prev = at
	v.next = p

	if p != nil {
		p.prev = v
	}

	l.len++
	return v
}

// 插入值到指定结点后
func (l *DList) insertVal(v interface{}, at *DListNode) *DListNode {
	return l.insert(&DListNode{Value: v}, at)
}

// Remove 移除结点, 返回移除结点的值
func (l *DList) Remove(v *DListNode) interface{} {
	if l.len == 0 {
		return nil
	}

	// 查找移除结点
	p := l.head
	for p != nil && p != v {
		p = p.next
	}

	if p == v {
		// 处理头结点
		if p == l.head {
			l.head = l.head.next
			l.head.prev = nil
		} else if p.next == nil {
			// 处理尾结点
			p.prev.next = nil
		} else {
			prev := p.prev
			next := p.next

			prev.next = next
			next.prev = prev
		}

		l.len--
	}

	return v.Value
}

// InsertBefore 在指定结点前插入值
func (l *DList) InsertBefore(v interface{}, mark *DListNode) *DListNode {
	return l.insertVal(v, mark.prev)
}

// InsertAfter 在指定结点后插入值
func (l *DList) InsertAfter(v interface{}, mark *DListNode) *DListNode {
	return l.insertVal(v, mark)
}

func (l *DList) String() string {
	s := fmt.Sprintf("Len: %d, Data: ", l.len)

	for p := l.head; p != nil; p = p.Next() {
		s += fmt.Sprintf("<-> %v ", p.Value)
	}
	return s
}
