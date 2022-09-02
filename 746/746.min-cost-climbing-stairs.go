package leetcode

/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	c0 := cost[0]
	c1 := cost[1]
	for i := 2; i < len(cost); i++ {
		if c0 < c1 {
			c0, c1 = c1, c0
			c1 = c1 + cost[i]
		} else {
			c0, c1 = c1, c1+cost[i]
		}
	}

	if c0 < c1 {
		return c0
	}
	return c1
}

// @lc code=end
