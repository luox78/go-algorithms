package sort_test

import (
	"testing"
)

func TestKthNum(t *testing.T) {
	// 编程实现 O(n) 时间复杂度内找到一组数据的第 K 大元素

	nums := []int{5, 2, 6, 1, 4, 3, 8}

	// 5
	t.Log(kthNum(nums, 5))

	// 4
	t.Log(kthNum(nums, 4))
}

// O(n) 时间复杂度内找到一组数据的第 K 大元素
func kthNum(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	// 第一次分区 O(n)
	q := partition(nums, 0, len(nums)-1)

	// O(n/2)+O(n/4)...
	for q+1 != k {
		if q+1 > k {
			// 在左区间
			q = partition(nums, 0, q-1)
		} else {
			// 在右区间
			q = partition(nums, q+1, len(nums)-1)
		}
	}

	return nums[q]
}

// 分区函数
func partition(nums []int, p, r int) int {
	if p >= r {
		return -1
	}

	// 分区点
	pivot := nums[r]

	i := p
	for j := p; j < r; j++ {
		if nums[j] < pivot {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp

			i++
		}
	}

	nums[r] = nums[i]
	nums[i] = pivot
	return i
}
