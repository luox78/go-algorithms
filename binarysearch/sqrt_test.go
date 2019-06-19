package binarysearch_test

import "testing"

func TestSqrt(t *testing.T) {
	// 计算并返回 x 的平方根，其中 x 是非负整数。

	// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

	// 示例 1:
	// 输入: 4
	// 输出: 2
	t.Log(sqrt(4))

	// 示例 2:

	// 输入: 8
	// 输出: 2
	// 说明: 8 的平方根是 2.82842...,
	// 	 由于返回类型是整数，小数部分将被舍去。
	t.Log(sqrt(8))

	t.Logf("%f", 10e-9)
}

func sqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	i, j := 1, x

	for i <= j {
		mid := i + (j-i)>>1

		if mid > x/mid {
			j = mid - 1
		} else if mid < x/mid {
			i = mid + 1
		} else {
			return mid
		}
	}

	return i - 1
}
