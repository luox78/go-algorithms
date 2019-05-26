//
// Package queue cyclequeue.go
// 循环队列实现
//
package queue

import "fmt"

// CycleQueue 循环队列结构
type CycleQueue struct {
	// 数据
	data []interface{}
	// 头指针
	head int
	// 尾指针
	tail int
	// 容量
	cap int
}

// NewCycleQueue 新建循环队列
func NewCycleQueue(cap int) *CycleQueue {
	if cap <= 0 {
		return nil
	}

	// 处理浪费的一个空间
	cap++

	return &CycleQueue{
		data: make([]interface{}, cap),
		cap:  cap,
	}
}

// Len 返回队列长度
func (q *CycleQueue) Len() int { return q.tail - q.head }

// Enqueue 入队
func (q *CycleQueue) Enqueue(v interface{}) bool {
	// 队列已满
	if (q.tail+1)%q.cap == q.head {
		return false
	}

	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.cap

	return true
}

// Dequeue 出队
func (q *CycleQueue) Dequeue() interface{} {
	// 队列已空
	if q.head == q.tail {
		return nil
	}

	v := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	return v
}

func (q *CycleQueue) String() string {
	return fmt.Sprintf("Len: %d, Data: %v", q.tail-q.head, q.data[q.head:q.tail])
}
