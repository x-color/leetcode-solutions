package leetcode

/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}
	n := 0
	c := 0
	for _, v := range m {
		n += v - v%2
		if v%2 == 1 {
			c = 1
		}
	}

	return n + c
}

// @lc code=end
