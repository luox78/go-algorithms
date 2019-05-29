package sort_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/sort"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{4, 6, 3, 2, 4, 7, 1}

	sort.BubbleSort(nums)

	// [1 2 3 4 4 6 7]
	t.Log(nums)
}
