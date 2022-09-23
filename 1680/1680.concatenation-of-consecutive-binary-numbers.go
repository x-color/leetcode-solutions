package leetcode

/*
 * @lc app=leetcode id=1680 lang=golang
 *
 * [1680] Concatenation of Consecutive Binary Numbers
 */

// @lc code=start
func concatenatedBinary(n int) int {
	var M int = 1e9 + 7
	ans := 0
	for i := 1; i <= n; i++ {
		m := 0
		for j := i; j > 0; j /= 2 {
			m++
		}
		ans = ((ans<<m)%M + i) % M
	}

	return ans
}

// @lc code=end
