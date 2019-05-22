//
// Package arrays dynamic.go
// 支持动态扩容数组的实现
//
package arrays

import "fmt"

// DynamicArray 动态数组结构
type DynamicArray struct {
	// 存放元素
	data []interface{}
	// 实际长度
	len int
	// 数组容量
	cap int
}

// NewDynamicArray 新建动态数组
func NewDynamicArray(cap int) *DynamicArray {
	return new(DynamicArray).init(cap)
}

// 初始化
func (a *DynamicArray) init(cap int) *DynamicArray {
	// 默认长度
	if cap <= 0 {
		cap = 4
	}

	a.data = make([]interface{}, cap)
	a.cap = cap
	a.len = 0
	return a
}

// Len 返回实际长度
func (a *DynamicArray) Len() int { return a.len }

// Cap 返回当前容量
func (a *DynamicArray) Cap() int { return a.cap }

// First 获取动态数组中第一个元素
func (a *DynamicArray) First() interface{} {
	return a.get(0)
}

// Last 获取动态数组中最后一个元素
func (a *DynamicArray) Last() interface{} {
	return a.get(a.len - 1)
}

// Get 根据索引获取元素
func (a *DynamicArray) Get(index int) interface{} {
	return a.get(index)
}

// Update 更新元素中的值
func (a *DynamicArray) Update(index int, val interface{}) bool {
	if index >= a.len || index < 0 {
		return false
	}

	a.data[index] = val
	return true
}

// InsertFirst 插入元素到数组第一个
func (a *DynamicArray) InsertFirst(val interface{}) {
	a.insert(0, val)
}

// InsertLast 插入元素到数组最后
func (a *DynamicArray) InsertLast(val interface{}) {
	a.insert(a.len, val)
}

// Insert 根据索引插入元素
func (a *DynamicArray) Insert(index int, val interface{}) bool {
	return a.insert(index, val)
}

// Remove 删除元素
func (a *DynamicArray) Remove(index int) bool {
	// 边界检查
	if index >= a.len || index < 0 {
		return false
	}

	// 数据搬移
	for i := index; i < a.len-1; i++ {
		a.data[i] = a.data[i+1]
	}

	a.len--

	if a.cap/2 >= a.len {
		a.resize()
	}
	return true
}

// 延迟初始化
func (a *DynamicArray) lazyInit() {
	if a.data == nil {
		a.init(0)
	}
}

// 调整容量
func (a *DynamicArray) resize() {
	// 缩容
	cap := a.cap / 2

	if a.len == a.cap {
		// 扩容
		cap = a.cap * 2
	}

	arr := make([]interface{}, cap)

	// 数据搬移
	copy(arr, a.data[:a.len])

	a.data = arr
	a.cap = cap
}

// 在索引处插入元素
func (a *DynamicArray) insert(index int, val interface{}) bool {
	a.lazyInit()

	// 边界检查
	if index > a.len || index < 0 {
		return false
	}

	if a.len == a.cap {
		a.resize()
	}

	// 数据搬移
	for i := a.len - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = val
	a.len++
	return true
}

// 根据索引获取元素
func (a *DynamicArray) get(index int) interface{} {
	a.lazyInit()

	// 边界检查
	if index >= a.len || index < 0 {
		return nil
	}

	return a.data[index]
}

func (a *DynamicArray) String() string {
	return fmt.Sprintf("Len: %d, Cap: %d, Data: %v", a.len, a.cap, a.data[:a.len])
}
