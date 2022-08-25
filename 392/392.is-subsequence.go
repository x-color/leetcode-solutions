package leetcode

import "strings"

/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	i := 0
	for _, r := range s {
		j := strings.IndexRune(t[i:], r)
		if j == -1 {
			return false
		}
		i += j + 1
	}

	return true
}

// @lc code=end
