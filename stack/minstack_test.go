package stack_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/stack"
)

func TestMinStack(t *testing.T) {
	ms := stack.NewMinStack()

	ms.Push(5)
	t.Log(ms.GetMin())

	ms.Push(4)
	t.Log(ms.GetMin())

	t.Log(ms.Pop())
	t.Log(ms.GetMin())
}
