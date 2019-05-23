//
// Package list slist.go
// 单链表实现
//
package list

import "fmt"

// SListNode 单链表节点
type SListNode struct {
	// 节点数据
	Value interface{}
	next  *SListNode
}

// Next 返回下一个节点
func (n *SListNode) Next() *SListNode {
	if n != nil {
		return n.next
	}
	return nil
}

// SList 单链表
type SList struct {
	// 头部节点
	head *SListNode
	// 单链表长度
	len int
}

// NewSList 新建单链表
func NewSList() *SList {
	return &SList{}
}

// Len 返回链表长度
func (l *SList) Len() int { return l.len }

// Front 返回第一个节点
func (l *SList) Front() *SListNode {
	if l.len == 0 {
		return nil
	}

	return l.head
}

// Back 返回最后一个节点
func (l *SList) Back() *SListNode {
	if l.len == 0 {
		return nil
	}

	p := l.head
	for p.next != nil {
		p = p.next
	}

	return p
}

// 在指定节点后插入元素
func (l *SList) insert(v, at *SListNode) *SListNode {
	n := at.next
	at.next = v
	v.next = n
	l.len++
	return v
}

// 在指定节点后插入值
func (l *SList) insertVal(v interface{}, at *SListNode) *SListNode {
	return l.insert(&SListNode{Value: v}, at)
}

// Remove 移除节点
func (l *SList) Remove(v *SListNode) interface{} {
	if v == nil || l.len == 0 {
		return nil
	}

	// 删除的是头节点
	if v == l.head {
		p := l.head
		l.head = l.head.next
		l.len--
		return p.Value
	}

	// 查找移除节点的前置节点
	p := l.head
	for p != nil && p.next != v {
		p = p.next
	}

	// 未找到匹配的节点
	if p == nil {
		return nil
	}

	p.next = p.next.next
	l.len--

	return v.Value
}

// PushFront 在头节点前加入元素
func (l *SList) PushFront(v interface{}) *SListNode {
	n := new(SListNode)
	n.Value = v
	n.next = l.head

	l.head = n
	l.len++
	return n
}

// PushBack 在尾节点后加入元素
func (l *SList) PushBack(v interface{}) *SListNode {
	p := l.head

	if p == nil {
		n := &SListNode{Value: v}
		l.head = n
		l.len++
		return n
	}

	for p.next != nil {
		p = p.next
	}

	return l.insertVal(v, p)
}

// InsertBefore 在指定节点之前插入元素
func (l *SList) InsertBefore(v interface{}, mark *SListNode) *SListNode {
	p := l.head
	for p != nil && p.next == mark {
		p = p.next
	}

	return l.insertVal(v, p)
}

// InsertAfter 在指定节点之后插入元素
func (l *SList) InsertAfter(v interface{}, mark *SListNode) *SListNode {
	return l.insertVal(v, mark)
}

func (l *SList) String() string {
	s := fmt.Sprintf("Len: %d, Data: ", l.len)

	for p := l.head; p != nil; p = p.Next() {
		s += fmt.Sprintf("-> %v ", p.Value)
	}
	return s
}
