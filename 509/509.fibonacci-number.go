package leetcode

/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(n int) int {
	n0 := 0
	n1 := 1
	if n == 0 {
		return n0
	}
	if n == 1 {
		return n1
	}
	for i := 2; i <= n; i++ {
		n0 = n0 + n1
		n0, n1 = n1, n0
	}

	return n1
}

// @lc code=end
