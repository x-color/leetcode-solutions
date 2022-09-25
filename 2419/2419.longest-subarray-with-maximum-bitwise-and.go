package leetcode

func longestSubarray(nums []int) int {
	c, ans, max := 0, 0, nums[0]
	for _, n := range nums {
		if n != max {
			c = 0
			if n > max {
				max = n
				ans = 0
			}
		}
		if n == max {
			c++
		}
		if c > ans {
			ans = c
		}
	}

	return ans
}
