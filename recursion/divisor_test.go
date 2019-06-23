package recursion_test

import "testing"

func TestCommonDivisor(t *testing.T) {
	// 找出最大公约数

	// 5
	t.Log(commonDivisor(25, 10))

	// 5
	t.Log(commonDivisor1(25, 10))

	// 5
	t.Log(commonDivisor2(25, 10))
}

// 辗转相除法, 最大公约数等于 min 和 max%min 余数之间的最大公约数
func commonDivisor(a, b int) int {
	max := a
	min := a
	if b > max {
		max = b
	}
	if b < min {
		min = b
	}
	if max%min == 0 {
		return min
	}

	return commonDivisor(min, max%min)
}

// 更相减损法, 最大公约数等于 max-min 和较小数 min 的最大公约数
func commonDivisor1(a, b int) int {
	max := a
	min := a
	if b > max {
		max = b
	}
	if b < min {
		min = b
	}
	if max%min == 0 {
		return min
	}

	return commonDivisor1(max-min, min)
}

func commonDivisor2(a, b int) int {
	if a == b {
		return a
	}
	// 当 a,b 为偶数时
	if a&1 == 0 && b&1 == 0 {
		return 2 * commonDivisor2(a>>1, b>>1)
	}
	// 当 a 为偶数, b 为奇数
	if a&1 == 0 && b&1 != 0 {
		return commonDivisor2(a>>1, b)
	}
	// 当 a 为奇数, b 为偶数
	if a&1 != 0 && b&1 == 0 {
		return commonDivisor2(a, b>>1)
	}
	// 当 a, b 为奇数, 先利用更相减损法运算, 此时 a-b 必然为偶数
	max := a
	min := a
	if b > max {
		max = b
	}
	if b < min {
		min = b
	}
	return commonDivisor2(max-min, min)
}
