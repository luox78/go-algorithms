package stack_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/stack"
)

func TestListStack(t *testing.T) {
	s := stack.NewListStack()

	// 入栈
	for i := 1; i <= 10; i++ {
		s.Push(i)
		t.Log(s)
	}

	// 出栈
	for i := 1; i <= 10; i++ {
		t.Log(s.Pop())
		t.Log(s)
	}

	t.Log(s.Pop())
	t.Log(s)
}
