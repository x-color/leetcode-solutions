package leetcode

func longestNiceSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	l := make([]int, len(nums))
	for i, n := range nums {
		l[i]++
		for _, m := range nums[i+1:] {
			if n&m != 0 {
				break
			}
			l[i]++
		}
	}

	max := 1
	for i, v := range l {
		lmax := v
		for _, m := range l[i+1:] {
			v--
			if v == 0 {
				break
			}
			if v > m {
				lmax -= v - m
				v = m
			}
		}
		if max < lmax {
			max = lmax
		}
	}
	return max
}
