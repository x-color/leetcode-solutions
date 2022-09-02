package leetcode

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	n0 := 0
	n1 := 1
	n2 := 2
	if n == 0 {
		return n0
	}
	if n == 1 {
		return n1
	}
	if n == 2 {
		return n2
	}

	for i := 3; i <= n; i++ {
		n1 = n1 + n2
		n1, n2 = n2, n1
	}

	return n2
}

// @lc code=end
