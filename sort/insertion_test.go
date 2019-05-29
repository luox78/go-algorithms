package sort_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/sort"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{2, 5, 1, 6, 8, 3}
	sort.InsertionSort(nums)

	// [1 2 3 5 6 8]
	t.Log(nums)
}
