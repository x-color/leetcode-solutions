package leetcode

func longestNiceSubarray(nums []int) int {
	j := 0
	used := 0
	max := 1
	for i := range nums {
		for (used & nums[i]) != 0 {
			used ^= nums[j]
			j++
		}
		used |= nums[i]
		if max < i-j+1 {
			max = i - j + 1
		}
	}
	return max
}
