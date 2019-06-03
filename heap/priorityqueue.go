//
// Package heap priorityqueue.go
// 实现优先级队列
//
package heap

import "fmt"

// PriorityQueue 优先级队列结构
type PriorityQueue struct {
	// 存储队列数据
	data []interface{}
	// 比较函数
	CompareFunc func(a, b interface{}) bool
	// 队列容量
	cap int
	// 队列长度
	len int
}

// NewPriorityQueue 新建优先级队列
func NewPriorityQueue(cap int, compareFunc func(a, b interface{}) bool) *PriorityQueue {
	if cap <= 0 || compareFunc == nil {
		return nil
	}

	return &PriorityQueue{
		data:        make([]interface{}, cap+1),
		CompareFunc: compareFunc,
		cap:         cap,
	}
}

// Len 返回队列长度
func (q *PriorityQueue) Len() int { return q.len }

// Cap 返回队列容量
func (q *PriorityQueue) Cap() int { return q.cap }

// Insert 插入元素
func (q *PriorityQueue) Insert(val interface{}) bool {
	if q.cap == q.len {
		return false
	}

	q.len++
	q.data[q.len] = val

	// 从下向上调整
	i := q.len
	for i/2 > 0 && q.CompareFunc(q.data[i], q.data[i/2]) {
		// 交换数据
		q.data[i], q.data[i/2] = q.data[i/2], q.data[i]
		i /= 2
	}

	return true
}

// Get 获取堆顶元素
func (q *PriorityQueue) Get() interface{} {
	if q.len == 0 {
		return nil
	}

	val := q.data[1]

	// 交换堆顶和最后元素
	q.data[1], q.data[q.len] = q.data[q.len], q.data[1]
	q.len--

	// 从上向下调整
	i := 1
	for {
		j := i
		// 和左子节点比较
		if i*2 <= q.len && q.CompareFunc(q.data[i*2], q.data[j]) {
			j = i * 2
		}
		// 和右子节点比较
		if i*2+1 <= q.len && q.CompareFunc(q.data[i*2+1], q.data[j]) {
			j = i*2 + 1
		}
		if j == i {
			break
		}

		// 交换数据
		q.data[i], q.data[j] = q.data[j], q.data[i]

		i = j
	}

	return val
}

func (q *PriorityQueue) String() string {
	return fmt.Sprintf("Cap: %d, Len: %d, Data: %v", q.cap, q.len, q.data[1:q.len+1])
}
