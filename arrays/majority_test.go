package arrays_test

import (
	"testing"
)

func TestMajority(t *testing.T) {
	// 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
	// 你可以假设数组是非空的，并且给定的数组总是存在众数。

	// 示例 1:
	// 输入: [3,2,3]
	// 输出: 3

	nums := []int{3, 2, 3}
	t.Log(majority1(nums))
	t.Log(majority2(nums))

	// 示例 2:
	// 输入: [2,2,1,1,1,2,2]
	// 输出: 2
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	t.Log(majority1(nums))
	t.Log(majority2(nums))
}

// 求众数, 散列表 解法
func majority1(nums []int) int {
	// 记录出现次数
	count := make(map[int]int)

	l := len(nums)
	for i := 0; i < l; i++ {
		count[nums[i]]++

		// 判断是否满足条件
		if count[nums[i]] > l/2 {
			return nums[i]
		}
	}
	return 0
}

// 求众数, 排序法, 因一定存在众数, 所以 n/2 的索引就是众数
func majority2(nums []int) int {
	qsort(nums, 0, len(nums)-1)
	return nums[len(nums)/2]
}
