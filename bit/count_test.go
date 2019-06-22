package bit_test

import "testing"

func TestHammingWeight(t *testing.T) {
	// 	编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。
	//
	// 	示例 1：
	// 	输入：00000000000000000000000000001011
	// 	输出：3
	// 	解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
	t.Log(hammingWeight(11))

	// 	示例 2：
	// 	输入：00000000000000000000000010000000
	// 	输出：1
	// 	解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。
	t.Log(hammingWeight(256))
}

// 统计二进制 1 的个数
func hammingWeight(num uint32) int {
	var count int

	for ; num != 0; num &= num - 1 {
		count++
	}
	return count
}
