package bit_test

import "testing"

func TestCountBits(t *testing.T) {
	// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

	// 示例 1:
	// 输入: 2
	// 输出: [0,1,1]
	t.Log(countBits(2))

	// 示例 2:
	// 输入: 5
	// 输出: [0,1,1,2,1,2]
	t.Log(countBits(5))
}

func countBits(num int) []int {
	res := make([]int, num+1)

	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}
