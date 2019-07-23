package lru

import "container/list"

// LRU LRU 结构
// 使用双向链表结合散列表实现的 LRU
// 时间复杂度 O(1)
type LRU struct {
	l   *list.List
	m   map[int]*list.Element
	len int
	cap int
}

// Data 缓冲数据结构
type Data struct {
	Key int
	Val int
}

// NewLRUData 根据 key, value 新建 LRUData
func NewLRUData(key, value int) *Data {
	return &Data{Key: key, Val: value}
}

// NewLRU 返回初始化 LRU
func NewLRU(cap int) *LRU {
	if cap <= 0 {
		return nil
	}

	return &LRU{
		l:   list.New(),
		m:   make(map[int]*list.Element, cap),
		cap: cap,
		len: 0,
	}
}

// Get 获取数据
func (lru *LRU) Get(key int) int {
	// 判断散列表是否存在
	if e, ok := lru.m[key]; ok {
		// 移动到最左边
		lru.l.MoveToFront(e)
		return e.Value.(*Data).Val
	}

	// 不存在
	return -1
}

// Put 设置数据
func (lru *LRU) Put(key int, value int) {
	// 判断是否存在, 存在则为更新
	if e, ok := lru.m[key]; ok {
		e.Value = NewLRUData(key, value)
		lru.l.MoveToFront(e)
		return
	}

	// 缓存满了
	if lru.len == lru.cap {
		// 删除最右边数据
		e := lru.l.Back()
		delete(lru.m, e.Value.(*Data).Key)
		lru.l.Remove(e)
		lru.len--
	}

	// 添加值最左边
	lru.l.PushFront(NewLRUData(key, value))
	lru.m[key] = lru.l.Front()
	lru.len++
}
