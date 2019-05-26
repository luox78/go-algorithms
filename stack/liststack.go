//
// Package stack liststack.go
// 链表实现的链式栈
//
package stack

import "fmt"

// ListStackNode 链表结点
type ListStackNode struct {
	// 数据
	Val interface{}
	// 下一个结点
	next *ListStackNode
}

// ListStack 链式栈结构
type ListStack struct {
	// 头结点
	head *ListStackNode
	// 尾结点
	tail *ListStackNode
	// 栈长度
	len int
}

// NewListStack 新建链式栈
func NewListStack() *ListStack {
	return new(ListStack)
}

// Len 返回栈长度
func (s *ListStack) Len() int { return s.len }

// Push 入栈
func (s *ListStack) Push(v interface{}) {
	if s.len == 0 {
		s.head = &ListStackNode{Val: v}
		s.tail = s.head
		s.len++
		return
	}

	s.tail.next = &ListStackNode{Val: v}
	s.tail = s.tail.next
	s.len++
}

// Pop 出栈
func (s *ListStack) Pop() interface{} {
	if s.len == 0 {
		return nil
	}
	n := s.head
	s.head = s.head.next
	s.len--
	if s.len == 0 {
		s.tail = nil
	}
	return n.Val
}

func (s *ListStack) String() string {
	res := fmt.Sprintf("Len: %d, ", s.len)

	p := s.head
	for p != nil {
		res += fmt.Sprintf("-> %v ", p.Val)
		p = p.next
	}

	return res
}
