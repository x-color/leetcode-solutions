package leetcode

var dp [1000][1001]int

func numberOfWays(startPos int, endPos int, k int) int {
	d := endPos - startPos
	if startPos > endPos {
		d = -d
	}
	return dfs(d, k)
}

func dfs(d, k int) int {
	if d < 0 {
		d = 1
	}
	if d == k {
		return 1
	}
	if d > k {
		return 0
	}
	if dp[d][k] == 0 {
		dp[d][k] = (1 + dfs(d+1, k-1) + dfs(d-1, k-1)) % (1e9 + 7)
	}
	return dp[d][k] - 1
}
