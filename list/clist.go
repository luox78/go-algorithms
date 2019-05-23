//
// Package list clist.go
// 循环链表实现
//
package list

import "fmt"

// CList 循环链表结构
type CList struct {
	// 数据
	Value interface{}
	// 下一个结点
	next *CList
}

// NewCList  新建循环链表
func NewCList() *CList {
	return new(CList)
}

// 初始化
func (l *CList) init(v interface{}) *CList {
	l.Value = v
	l.next = l
	return l
}

// Next 下一个结点
func (l *CList) Next() *CList {
	if l.next != nil {
		return l.next
	}
	return nil
}

// Insert 插入结点值
func (l *CList) Insert(v interface{}) *CList {
	if l.next == nil {
		return l.init(v)
	}

	p := l.next
	l.next = &CList{Value: v, next: p}

	return l.next
}

// Update 更新结点值
func (l *CList) Update(v interface{}) *CList {
	l.Value = v
	return l
}

// Remove 移除结点, 返回移除结点的值
func (l *CList) Remove() interface{} {
	if l.next == nil {
		return nil
	}

	// 找到移除结点的前驱结点
	p := l
	for ; p.next != l; p = p.next {
	}

	r := p.next

	p.next = p.next.next

	r.next = nil
	return r.Value
}

// Len 返回循环链表长度
func (l *CList) Len() int {
	if l.next == nil {
		return 0
	}

	len := 1
	for p := l.next; p != l; p = p.next {
		len++
	}
	return len
}

func (l *CList) String() string {
	var s string
	if l.next != nil {
		s += fmt.Sprintf("-> %v ", l.Value)
		for p := l.next; p != l; p = p.next {
			s += fmt.Sprintf("-> %v ", p.Value)
		}
	}
	return s
}
