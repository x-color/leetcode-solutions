package leetcode

/*
 * @lc app=leetcode id=1578 lang=golang
 *
 * [1578] Minimum Time to Make Rope Colorful
 */

// @lc code=start
func minCost(colors string, neededTime []int) int {
	ans := 0
	for i := 0; i < len(colors); i++ {
		c := colors[i]
		max := neededTime[i]
		sum := 0
		for i < len(colors) && c == colors[i] {
			sum += neededTime[i]
			if neededTime[i] > max {
				max = neededTime[i]
			}
			i++
		}

		if sum > 0 {
			ans += sum - max
			i--
		}
	}

	return ans
}

// @lc code=end
