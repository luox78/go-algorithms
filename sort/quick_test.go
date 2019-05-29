package sort_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/sort"
)

func TestQuickSort(t *testing.T) {
	nums := []int{6, 3, 8, 1, 7, 2, 5, 4}
	sort.QuickSort(nums)

	// [1 2 3 4 5 6 7 8]
	t.Log(nums)
}
