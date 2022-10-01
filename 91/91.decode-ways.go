package leetcode

/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}

	for i := 1; i < len(s); i++ {
		j := i + 1
		if s[i] != '0' {
			dp[j] += dp[j-1]
		}
		if s[i-1] == '1' || s[i-1] == '2' && s[i]-'0' < 7 {
			dp[j] += dp[j-2]
		}
	}

	return dp[len(s)]
}

// @lc code=end
