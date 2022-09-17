package leetcode

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	record := map[int]bool{}
	for {
		record[n] = true
		s := 0
		for n > 0 {
			m := n % 10
			n = n / 10
			s += m * m
		}
		if s == 1 {
			return true
		}
		n = s
		if record[n] {
			return false
		}
	}
}

// @lc code=end
