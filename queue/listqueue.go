//
// Package queue listqueue.go
// 链表实现的链式队列
//
package queue

import "fmt"

// ListQueueNode 链表结点
type ListQueueNode struct {
	// 数据
	Val interface{}
	// 下一个结点
	next *ListQueueNode
}

// ListQueue 链式队列结构
type ListQueue struct {
	// 头结点
	head *ListQueueNode
	// 尾结点
	tail *ListQueueNode
	// 队列长度
	len int
}

// NewListQueue 新建链式队列
func NewListQueue() *ListQueue {
	return new(ListQueue)
}

// Len 返回链式队列长度
func (q *ListQueue) Len() int { return q.len }

// Enqueue 入队
func (q *ListQueue) Enqueue(v interface{}) {
	if q.len == 0 {
		q.head = &ListQueueNode{Val: v}
		q.tail = q.head
		q.len++
		return
	}

	q.tail.next = &ListQueueNode{Val: v}
	q.tail = q.tail.next
	q.len++
}

// Dequeue 出队
func (q *ListQueue) Dequeue() interface{} {
	if q.len == 0 {
		return nil
	}

	n := q.head
	q.head = q.head.next
	q.len--
	return n.Val
}

func (q *ListQueue) String() string {
	res := fmt.Sprintf("Len: %d,", q.len)
	p := q.head
	for p != nil {
		res += fmt.Sprintf("-> %v ", p.Val)
		p = p.next
	}
	return res
}
