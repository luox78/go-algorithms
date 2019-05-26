//
// Package stack arraystack.go
// 数组实现的顺序栈
//
package stack

import "fmt"

// ArrayStack 顺序栈结构
type ArrayStack struct {
	// 数据
	data []interface{}
	// 栈长度
	len int
	// 栈容量
	cap int
}

// NewArrayStack 新建顺序栈
func NewArrayStack(cap int) *ArrayStack {
	if cap <= 0 {
		return nil
	}

	return &ArrayStack{
		data: make([]interface{}, cap),
		cap:  cap,
	}
}

// Len 返回栈长度
func (s *ArrayStack) Len() int { return s.len }

// Cap 返回栈容量
func (s *ArrayStack) Cap() int { return s.cap }

// Push 入栈
func (s *ArrayStack) Push(v interface{}) bool {
	// 栈已满
	if s.len == s.cap {
		return false
	}

	s.data[s.len] = v
	s.len++

	return true
}

// Pop 出栈
func (s *ArrayStack) Pop() interface{} {
	// 栈已空
	if s.len == 0 {
		return nil
	}

	v := s.data[s.len-1]

	s.len--
	return v
}

func (s *ArrayStack) String() string {
	return fmt.Sprintf("Len: %d, Cap: %d, Data: %v", s.len, s.cap, s.data[:s.len])
}
