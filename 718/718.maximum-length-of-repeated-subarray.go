package leetcode

/*
 * @lc app=leetcode id=718 lang=golang
 *
 * [718] Maximum Length of Repeated Subarray
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	ans := 0
	memo := make([][]int, len(nums1)+1)
	for i := range memo {
		memo[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				memo[i][j] = memo[i-1][j-1] + 1
				if ans < memo[i][j] {
					ans = memo[i][j]
				}
			}
		}
	}

	return ans
}

// @lc code=end
