package leetcode

/*
 * @lc app=leetcode id=990 lang=golang
 *
 * [990] Satisfiability of Equality Equations
 */

// @lc code=start
func equationsPossible(equations []string) bool {
	uf := [26]int{}
	for i := range uf {
		uf[i] = i
	}
	for _, eq := range equations {
		if eq[1] == '=' {
			uf[parent(uf, int(eq[0]-'a'))] = parent(uf, int(eq[3]-'a'))
		}
	}
	for _, eq := range equations {
		if eq[1] == '!' && parent(uf, int(eq[0]-'a')) == parent(uf, int(eq[3]-'a')) {
			return false
		}
	}
	return true
}

func parent(uf [26]int, x int) int {
	if uf[x] == x {
		return x
	}
	uf[x] = parent(uf, uf[x])
	return uf[x]
}

// @lc code=end
