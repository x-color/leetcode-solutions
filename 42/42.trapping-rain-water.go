package leetcode

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	maxLeft, maxRight := height[l], height[r]
	for l <= r {
		if maxLeft < maxRight {
			if height[l] > maxLeft {
				maxLeft = height[l]
			}
			ans += maxLeft - height[l]
			l++
		} else {
			if height[r] > maxRight {
				maxRight = height[r]
			}
			ans += maxRight - height[r]
			r--
		}
	}

	return ans
}

// @lc code=end
