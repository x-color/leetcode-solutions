package leetcode

/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 */

// @lc code=start
func pivotIndex(nums []int) int {
	sr := 0
	for _, n := range nums {
		sr += n
	}

	sl := 0
	for i, n := range nums {
		sr -= n
		if sl == sr {
			return i
		}
		sl += n
	}
	return -1
}

// @lc code=end
