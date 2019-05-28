package recursion_test

import (
	"fmt"
	"testing"
)

func TestPermutation(t *testing.T) {
	// [1 2 3]
	// [1 3 2]
	// [2 1 3]
	// [2 3 1]
	// [3 1 2]
	// [3 2 1]
	n := []int{1, 2, 3}
	var res []int
	permutate(n, res)
}

// 求全排列
func permutate(nums []int, res []int) {
	if len(nums) == 0 {
		fmt.Println(res)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 选择一个数字加入结果
		newRes := make([]int, len(res)+1)
		copy(newRes, res)
		newRes[len(newRes)-1] = nums[i]

		// 移除选择的数字继续后面的选择
		newNums := make([]int, len(nums)-1)
		copy(newNums[:i], nums[:i])
		copy(newNums[i:], nums[i+1:])

		permutate(newNums, newRes)
	}
}
