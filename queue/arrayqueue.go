//
// Package queue arrayqueue.go
// 数组实现的顺序队列
//
package queue

import "fmt"

// ArrayQueue 顺序队列结构
type ArrayQueue struct {
	// 数据
	data []interface{}
	// 队列长度
	len int
	// 队列容量
	cap int
}

// NewArrayQueue 新建顺序队列
func NewArrayQueue(cap int) *ArrayQueue {
	if cap <= 0 {
		return nil
	}
	return &ArrayQueue{
		data: make([]interface{}, cap),
		cap:  cap,
	}
}

// Len 返回队列长度
func (q *ArrayQueue) Len() int { return q.len }

// Cap 返回队列容量
func (q *ArrayQueue) Cap() int { return q.cap }

// Enqueue 入队
func (q *ArrayQueue) Enqueue(v interface{}) bool {
	// 队列已满
	if q.len == q.cap {
		return false
	}

	q.data[q.len] = v
	q.len++

	return true
}

// Dequeue 出队
func (q *ArrayQueue) Dequeue() interface{} {
	// 队列已空
	if q.len == 0 {
		return nil
	}

	v := q.data[0]

	// 数据搬移
	for i := 0; i < q.len-1; i++ {
		q.data[i] = q.data[i+1]
	}

	q.len--
	return v
}

func (q *ArrayQueue) String() string {
	return fmt.Sprintf("Len: %d,Cap: %d, Data: %v", q.len, q.cap, q.data[:q.len])
}
