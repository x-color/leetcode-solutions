package leetcode

import "sort"

func answerQueries(nums []int, queries []int) []int {
	answer := make([]int, len(queries))
	sort.Ints(nums)

	for i, q := range queries {
		s := 0
		for _, n := range nums {
			s += n
			answer[i]++
			if s > q {
				answer[i]--
				break
			}
		}
	}

	return answer
}

// 100, 1, 1 nums
// q =100
// -> 2
