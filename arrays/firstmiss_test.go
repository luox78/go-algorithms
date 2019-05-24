package arrays_test

import (
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	// 给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
	// 说明: 算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。

	// 示例 1:
	// 输入: [1,2,0]
	// 输出: 3
	nums := []int{1, 2, 0}
	t.Log(firstMissingPositive(nums))

	// 示例 2:
	// 输入: [3,4,-1,1]
	// 输出: 2
	nums = []int{3, 4, -1, 1}
	t.Log(firstMissingPositive(nums))

	// 示例 3:
	// 输入: [7,8,9,11,12]
	// 输出: 1
	nums = []int{7, 8, 9, 11, 12}
	t.Log(firstMissingPositive(nums))

}

// 求第一个缺失的正整数, 时间复杂度 O(n) , 空间复杂度 O(1)
func firstMissingPositive(nums []int) int {
	n := len(nums)

	// 判断是否存在 1
	var hasOne bool
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			hasOne = true
		}
	}

	// 不存在数字 1, 返回 1
	if !hasOne {
		return 1
	}

	// 只有一个数字 [1]
	if n == 1 {
		return 2
	}

	// 负数、0 和大于 n 的数字不考虑, 直接替换为 1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}

	// 以数组索引代替数字作为 key, 正负值作为判断数字是否存在
	for i := 0; i < n; i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}

		if num == n {
			// 等于 n , 用索引 0 存放
			if nums[0] > 0 {
				// 避免重复修改
				nums[0] = -nums[0]
			}
			continue
		}

		// 避免重复修改
		if nums[num] > 0 {
			nums[num] = -nums[num]
		}
	}

	// 判断数组中正数出现的索引, 该索引为缺失的正数
	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	// 小于 n 的数字都已存在, 判断 n 是否为缺失的
	if nums[0] > 0 {
		return n
	}

	// 小于等于 n 的数字都已存在, 缺失的正数为 n+1
	return n + 1
}
