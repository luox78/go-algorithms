package string

import "testing"

func TestReverseString(t *testing.T) {
	// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

	// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

	// 你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

	// 示例 1：
	// 输入：["h","e","l","l","o"]
	// 输出：["o","l","l","e","h"]
	s := []byte("he")
	reverseString1(s)
	t.Log(string(s))

	// 示例 2：
	// 输入：["H","a","n","n","a","h"]
	// 输出：["h","a","n","n","a","H"]
	s = []byte("Hannah")
	reverseString1(s)
	t.Log(string(s))
}

func reverseString(s []byte) {
	n := len(s)
	if n < 2 {
		return
	}

	for i, l := 0, n/2; i < l; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}

func reverseString1(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
