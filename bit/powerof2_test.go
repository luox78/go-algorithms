package bit_test

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	// 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

	// 示例 1:
	// 输入: 1
	// 输出: true
	// 解释: 2^0 = 1
	t.Log(isPowerOfTwo(1))

	// 示例 2:
	// 输入: 16
	// 输出: true
	// 解释: 2^ = 16
	t.Log(isPowerOfTwo(16))

	// 示例 3:
	// 输入: 218
	// 输出: false
	t.Log(isPowerOfTwo(218))
}

// 判断n是否是 2 的幂次方
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
