package stack_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/stack"
)

func TestArrayStack(t *testing.T) {
	s := stack.NewArrayStack(10)

	// 入栈
	for i := 1; i <= 10; i++ {
		s.Push(i)
		t.Log(s)
	}

	s.Push(11)
	t.Log(s)

	// 出栈
	for i := 1; i <= 10; i++ {
		t.Log(s.Pop())
	}

	t.Log(s.Pop())
}
