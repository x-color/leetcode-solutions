package leetcode

/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	ans := 0
	count := map[rune]int{}
	i := 0
	max_count := 0
	for j, r := range s {
		count[r]++
		if max_count < count[r] {
			max_count = count[r]
		}
		length := j - i + 1
		if length-max_count > k {
			count[rune(s[i])]--
			i++
			continue
		}

		if ans < length {
			ans = length
		}
	}

	return ans
}

// @lc code=end
