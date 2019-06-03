package heap_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/heap"
)

func TestPriorityQueue(t *testing.T) {
	q := heap.NewPriorityQueue(10, compareFunc)

	for i := 1; i <= 10; i++ {
		q.Insert(i)
	}

	// Cap: 10, Len: 10, Data: [10 9 6 7 8 2 5 1 4 3]
	t.Log(q)

	for q.Len() > 0 {
		t.Log(q.Get())
	}
}

func compareFunc(a, b interface{}) bool {
	return a.(int) > b.(int)
}
