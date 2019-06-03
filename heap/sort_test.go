package heap_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/heap"
)

func TestHeapSort(t *testing.T) {
	nums := []int{6, 3, 7, 3, 8, 2, 1, 9}

	heap.Sort(nums)

	// [1 2 3 3 6 7 8 9]
	t.Log(nums)
}
