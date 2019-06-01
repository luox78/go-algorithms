//
// Package hashtable listhashtable.go
// 实现基于链表法解决冲突的散列表
//
package hashtable

import "hash/crc32"

// listNode 链表
type listNode struct {
	value *data
	next  *listNode
}

// 散列表数据
type data struct {
	key string
	val interface{}
}

// ListHashTable 链表法散列表结构
type ListHashTable struct {
	data []*listNode
	len  int
}

// NewListHashTable 返回新建的散列表
func NewListHashTable(l int) *ListHashTable {
	if l > 0 {
		return &ListHashTable{
			data: make([]*listNode, l),
			len:  l,
		}
	}

	return nil
}

// 返回新建的存储数据
func (h *ListHashTable) initData(key string, val interface{}) *data {
	return &data{key: key, val: val}
}

// Set 存储数据
func (h *ListHashTable) Set(key string, val interface{}) {
	// 获取索引
	index := h.hashKey(key)

	// 获取头指针
	p := h.data[index]
	if p == nil {
		h.data[index] = &listNode{value: h.initData(key, val)}
		return
	}

	for p.next != nil {
		if p.value.key == key {
			// 修改数据
			p.value = h.initData(key, val)
			return
		}
		p = p.next
	}

	// 新增数据
	p.next = &listNode{value: h.initData(key, val)}
}

// Get 获取数据
func (h *ListHashTable) Get(key string) interface{} {
	// 获取索引
	index := h.hashKey(key)

	// 获取头指针
	p := h.data[index]

	for p != nil {
		if p.value.key == key {
			return p.value.val
		}
		p = p.next
	}

	return nil
}

// 哈希函数
func (h *ListHashTable) hashKey(key string) int {
	return int(crc32.ChecksumIEEE([]byte(key))) % h.len
}
