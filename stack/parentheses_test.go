package stack_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/stack"
)

func TestVaildParentheses(t *testing.T) {
	// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

	// 有效字符串需满足：

	// 左括号必须用相同类型的右括号闭合。
	// 左括号必须以正确的顺序闭合。
	// 注意空字符串可被认为是有效字符串。

	// 示例 1:

	// 输入: "()"
	// 输出: true
	t.Log(vaildParentheses("()"))

	// 示例 2:
	// 输入: "()[]{}"
	// 输出: true
	t.Log(vaildParentheses("()[]{}"))

	// 示例 3:
	// 输入: "(]"
	// 输出: false
	t.Log(vaildParentheses("(]"))

	// 示例 4:
	// 输入: "([)]"
	// 输出: false
	t.Log(vaildParentheses("([)]"))

	// 示例 5:
	// 输入: "{[]}"
	// 输出: true
	t.Log(vaildParentheses("{[]}"))
}

// 有效的括号
func vaildParentheses(str string) bool {
	// 记录括号匹配映射
	mappings := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	// 利用栈特性进行匹配
	stack := stack.NewArrayStack(len(str))

	for _, c := range str {
		s := string(c)

		if ms, ok := mappings[s]; !ok {
			// 未找到映射, 存入栈
			stack.Push(s)
		} else {
			// 从栈中取出一个字符串进行比较
			if ss := stack.Pop(); ms != ss {
				return false
			}
		}

	}

	return true
}
