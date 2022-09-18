package leetcode

import "sort"

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i int, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	ans := ""
	for i := 0; i < len(strs[0]); i++ {
		b := strs[0][i]
		for _, s := range strs {
			if s[i] != b {
				return ans
			}
		}
		ans += string(b)
	}

	return ans
}

// @lc code=end
