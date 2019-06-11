package recursion_test

import "testing"

func TestMyPow(t *testing.T) {
	// 实现 pow(x, n) ，即计算 x 的 n 次幂函数。
	// 说明:
	// -100.0 < x < 100.0
	// n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

	// 示例 1:
	// 输入: 2.00000, 10
	// 输出: 1024.00000
	t.Log(myPow1(2.0, 10))

	// 示例 2:
	// 输入: 2.10000, 3
	// 输出: 9.26100
	t.Log(myPow1(2.1, 3))

	// 示例 3:
	// 输入: 2.00000, -2
	// 输出: 0.25000
	// 解释: 2-2 = 1/22 = 1/4 = 0.25
	t.Log(myPow1(2.0, -2))
}

// 计算 x 的 n 次幂, 非递归实现
func myPow1(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	pow := 1.0

	for n > 0 {
		if n&1 == 1 {
			pow *= x
		}
		x *= x
		n = n >> 1
	}
	return pow
}
