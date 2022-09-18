package leetcode

func longestContinuousSubstring(s string) int {
	l := 0
	max := 1
	for r := 1; r < len(s); r++ {
		if s[r]-s[r-1] != 1 {
			l = r
		}
		if max < r-l+1 {
			max = r - l + 1
		}
	}
	return max
}
