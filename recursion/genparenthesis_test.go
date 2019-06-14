package recursion_test

import "testing"

func TestGenParenthesis(t *testing.T) {
	// 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

	// 例如，给出 n = 3，生成结果为：
	// [
	//   "((()))",
	//   "(()())",
	//   "(())()",
	//   "()(())",
	//   "()()()"
	// ]

	t.Log(genParenthesis(3))
}

// 生成 n 个有效的括号组合
func genParenthesis(n int) []string {
	var result []string

	gen(&result, "", n, 0, 0)

	return result
}

// 递归剪枝
func gen(result *[]string, sub string, n, left, right int) {
	// 已满足条件
	if left == n && right == n {
		*result = append(*result, sub)
		return
	}

	// 判断生成左括号条件
	if left < n {
		gen(result, sub+"(", n, left+1, right)
	}

	// 判断生成右括号并且有效的条件
	if left > right && right < n {
		gen(result, sub+")", n, left, right+1)
	}
}
