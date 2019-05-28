package recursion_test

import "testing"

func TestFib(t *testing.T) {
	// 55
	t.Log(fib(10))

	// 6765
	t.Log(fib(20))

	// 102334155
	t.Log(fib(40))
}

// 斐波那契数列求值
func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
