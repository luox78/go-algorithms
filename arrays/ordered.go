//
// Package arrays ordered.go
// 大小固定 支持动态操作的有序数组实现
//
package arrays

import (
	"errors"
	"fmt"
)

// OrderedArray 有序数组结构
type OrderedArray struct {
	// 数组数据
	data []int
	// 数组实际长度
	len int
	// 数组容量
	cap int
}

// NewOrderedArray 新建有序数组
func NewOrderedArray(cap int) *OrderedArray {
	return &OrderedArray{
		data: make([]int, cap),
		cap:  cap,
	}
}

// Add 新增
func (a *OrderedArray) Add(val int) int {
	// 有序数组容量满了
	if a.len == a.cap {
		return -1
	}

	// 查找插入的位置
	i := a.len - 1
	for ; i >= 0; i-- {
		if a.data[i] > val {
			// 数据搬移
			a.data[i+1] = a.data[i]
		} else {
			break
		}
	}

	index := i + 1
	a.data[index] = val
	a.len++

	return index
}

// Del 删除
func (a *OrderedArray) Del(index int) {
	// 边界检查
	if index >= a.len || index < 0 {
		return
	}

	// 数据搬移
	for i := index; i < a.len-1; i++ {
		a.data[i] = a.data[i+1]
	}

	a.len--
}

// Update 更新值
func (a *OrderedArray) Update(index, val int) int {
	// 边界检查
	if index >= a.len || index < 0 {
		return -1
	}

	a.Del(index)
	return a.Add(val)
}

// Get 读取
func (a *OrderedArray) Get(index int) (int, error) {
	// 边界检查
	if index >= a.len || index < 0 {
		return 0, errors.New("无效索引")
	}

	return a.data[index], nil
}

// Len 返回有序数组长度
func (a *OrderedArray) Len() int { return a.len }

// Cap 返回有序数组容量
func (a *OrderedArray) Cap() int { return a.cap }

func (a *OrderedArray) String() string {
	return fmt.Sprintf("Len: %d, Cap: %d, Data: %v", a.len, a.cap, a.data[:a.len])
}
