package arrays_test

import "testing"

func TestProductExceptSelf(t *testing.T) {
	// 给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

	// 示例:
	// 输入: [1,2,3,4]
	// 输出: [24,12,8,6]
	// 说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
	t.Log(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	left, right := 1, 1
	n := len(nums)
	res := make([]int, n)

	// 计算左边除自身之外的乘积
	for i := 0; i < n; i++ {
		res[i] = left
		left *= nums[i]
	}
	// 计算右边除自身之外的乘积
	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}
