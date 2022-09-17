package leetcode

import "sort"

/*
 * @lc app=leetcode id=948 lang=golang
 *
 * [948] Bag of Tokens
 */

// @lc code=start
func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)

	score := 0
	result := 0

	for len(tokens) > 0 {
		if power >= tokens[0] {
			power -= tokens[0]
			score++

			if score > result {
				result = score
			}

			tokens = tokens[1:]
			continue
		}

		if score == 0 {
			return result
		}

		power += tokens[len(tokens)-1]
		score--
		tokens = tokens[:len(tokens)-1]
	}

	return result
}

// @lc code=end
