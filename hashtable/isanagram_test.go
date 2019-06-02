package hashtable_test

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

	// 示例 1:
	// 输入: s = "anagram", t = "nagaram"
	// 输出: true
	t.Log(isAnagram("anagram", "nagaram"))

	// 示例 2:
	// 输入: s = "rat", t = "car"
	// 输出: false
	t.Log(isAnagram("rat", "car"))

	// 说明:
	// 你可以假设字符串只包含小写字母。

	t.Log(isAnagram("ab", "a"))
}

func isAnagram(s string, t string) bool {
	m1 := make(map[rune]int)

	for _, c := range s {
		m1[c] = m1[c] + 1
	}

	for _, c := range t {
		if count, ok := m1[c]; !ok {
			return false
		} else {
			if count == 1 {
				delete(m1, c)
			} else {
				m1[c] = m1[c] - 1
			}
		}
	}

	return len(m1) == 0
}
