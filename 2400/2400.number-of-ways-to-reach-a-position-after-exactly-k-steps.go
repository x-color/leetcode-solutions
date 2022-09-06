package leetcode

import (
	"fmt"
)

var MOD int = 1e9 + 7

func numberOfWays(startPos int, endPos int, k int) int {
	if (k-endPos+startPos)%2 != 0 {
		return 0
	}
	memo := map[string]int{}
	return dfs(startPos, k, endPos, memo) % MOD
}

func dfs(pos, k, endPos int, memo map[string]int) int {
	if k == 0 && pos == endPos {
		return 1
	}
	if k == 0 || pos-endPos > k || endPos-pos > k {
		return 0
	}
	key := fmt.Sprintf("%d,%d", pos, k)
	if v, ok := memo[key]; ok {
		return v
	}

	memo[key] = (dfs(pos-1, k-1, endPos, memo)%MOD + dfs(pos+1, k-1, endPos, memo)%MOD) % MOD

	return memo[key]
}
