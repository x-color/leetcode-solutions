package leetcode

/*
 * @lc app=leetcode id=658 lang=golang
 *
 * [658] Find K Closest Elements
 */

// @lc code=start
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	for r-l > k {
		if abs(arr[l], x) > abs(arr[r], x) {
			l++
		} else {
			r--
		}
	}

	ans := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		ans = append(ans, arr[i])
	}

	return ans
}

// @lc code=end
