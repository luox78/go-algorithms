package list_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/list"
)

func TestCList(t *testing.T) {
	l := list.NewCList()

	// -> 1 -> 2 -> 3
	n := l.Insert(1).Insert(2).Insert(3)
	t.Log(l)

	// -> 1 -> 2 -> 4
	n.Update(4)
	t.Log(l)

	// 1
	t.Log(n.Next().Value)

	// 3
	t.Log(l.Len())

	// 4
	t.Log(n.Remove())

	// 2 | -> 1 -> 2
	t.Log(l.Len(), "|", l)
}
