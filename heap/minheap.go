//
// Package heap minheap.go
// 小顶堆实现
//
package heap

import "fmt"

// MinHeap 小顶堆结构
type MinHeap struct {
	data []int
	cap  int
	len  int
}

// NewMinHeap 新建小顶堆
func NewMinHeap(cap int) *MinHeap {
	if cap <= 0 {
		return nil
	}

	// 索引 0 不存放数据
	return &MinHeap{
		data: make([]int, cap+1),
		cap:  cap,
	}
}

// Len 返回小顶堆长度
func (h *MinHeap) Len() int { return h.len }

// Cap 返回小顶堆容量
func (h *MinHeap) Cap() int { return h.cap }

// Insert 插入元素
func (h *MinHeap) Insert(val int) bool {
	if h.len == h.cap {
		return false
	}

	h.len++
	h.data[h.len] = val

	// 堆化, 从下向上调整
	i := h.len
	for i/2 > 0 && h.data[i/2] > h.data[i] {
		// 交换数据
		h.data[i/2], h.data[i] = h.data[i], h.data[i/2]
		i /= 2
	}

	return true
}

// RemoveMin 移除堆顶元素
func (h *MinHeap) RemoveMin() int {
	if h.len == 0 {
		return -1
	}

	// 获取堆顶元素
	val := h.data[1]

	// 顶堆和最后元素交换
	h.data[1], h.data[h.len] = h.data[h.len], h.data[1]

	h.len--

	// 从上向下调整
	i := 1
	for {
		minPos := i
		// 判断是否比左子节点小
		if i*2 <= h.len && h.data[minPos] > h.data[i*2] {
			minPos = i * 2
		}
		// 判断是否比右子节点小
		if i*2+1 <= h.len && h.data[minPos] > h.data[i*2+1] {
			minPos = i*2 + 1
		}
		if i == minPos {
			break
		}
		// 交换数据
		h.data[i], h.data[minPos] = h.data[minPos], h.data[i]
		i = minPos
	}

	return val
}

func (h *MinHeap) String() string {
	return fmt.Sprintf("Cap: %d, Len: %d, Data: %v", h.cap, h.len, h.data[1:h.len+1])
}
