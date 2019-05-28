package recursion_test

import "testing"

func TestFactorial(t *testing.T) {
	// 120
	t.Log(factorial(5))

	// 3628800
	t.Log(factorial(10))
}

// 求阶乘
func factorial(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
