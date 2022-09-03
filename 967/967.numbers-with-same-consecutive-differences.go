package leetcode

/*
 * @lc app=leetcode id=967 lang=golang
 *
 * [967] Numbers With Same Consecutive Differences
 */

// @lc code=start
func numsSameConsecDiff(n int, k int) []int {
	ans := []int{}
	for i := 1; i <= 9; i++ {
		dfs(i, k, n-1, &ans)
	}

	return ans
}

func dfs(n, k, d int, ans *[]int) {
	if d == 0 {
		*ans = append(*ans, n)
		return
	}

	m := n % 10
	if m-k >= 0 {
		dfs(n*10+m-k, k, d-1, ans)
	}
	if m+k <= 9 && k != 0 {
		dfs(n*10+m+k, k, d-1, ans)
	}
}

// @lc code=end
