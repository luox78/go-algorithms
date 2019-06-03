//
// Package heap maxheap.go
// 大顶堆实现
//
package heap

import "fmt"

// MaxHeap 大顶堆结构
type MaxHeap struct {
	data []int
	cap  int
	len  int
}

// NewMaxHeap 新建大顶堆
func NewMaxHeap(cap int) *MaxHeap {
	if cap <= 0 {
		return nil
	}

	return &MaxHeap{
		data: make([]int, cap+1),
		cap:  cap,
	}
}

// Len 返回大顶堆长度
func (h *MaxHeap) Len() int { return h.len }

// Cap 返回大顶堆容量
func (h *MaxHeap) Cap() int { return h.cap }

// Insert 插入元素
func (h *MaxHeap) Insert(val int) bool {
	if h.len == h.cap {
		return false
	}

	h.len++
	h.data[h.len] = val

	// 从下向上调整
	i := h.len
	for i/2 > 0 && h.data[i/2] < h.data[i] {
		// 交换数据
		h.data[i/2], h.data[i] = h.data[i], h.data[i/2]
		i /= 2
	}

	return true
}

// RemoveMax 移除堆顶元素
func (h *MaxHeap) RemoveMax() int {
	if h.len == 0 {
		return -1
	}

	val := h.data[1]

	// 顶堆元素和最后元素交换
	h.data[1], h.data[h.len] = h.data[h.len], h.data[1]
	h.len--

	// 从上向下调整
	i := 1
	for {
		maxPos := i
		// 和左子节点比较
		if i*2 <= h.len && h.data[maxPos] < h.data[i*2] {
			maxPos = i * 2
		}
		// 和右子节点比较
		if i*2+1 <= h.len && h.data[maxPos] < h.data[i*2+1] {
			maxPos = i*2 + 1
		}

		if maxPos == i {
			break
		}
		// 交换数据
		h.data[i], h.data[maxPos] = h.data[maxPos], h.data[i]

		i = maxPos
	}

	return val
}

func (h *MaxHeap) String() string {
	return fmt.Sprintf("Cap: %d, Len: %d, Data: %v", h.cap, h.len, h.data[1:h.len+1])
}
