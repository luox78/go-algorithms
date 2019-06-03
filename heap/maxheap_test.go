package heap_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/heap"
)

func TestMaxHeap(t *testing.T) {
	h := heap.NewMaxHeap(10)

	for i := 1; i <= 10; i++ {
		h.Insert(i)
	}

	// Cap: 10, Len: 10, Data: [1 2 5 4 3 9 6 10 7 8]
	t.Log(h)

	for h.Len() > 0 {
		t.Log(h.RemoveMax())
		t.Log(h)
	}
}
