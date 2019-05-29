package sort_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/sort"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{6, 3, 3, 5, 1, 8, 2, 9}
	sort.SelectionSort(nums)

	// [1 2 3 3 5 6 8 9]
	t.Log(nums)
}
