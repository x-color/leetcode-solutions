package leetcode

func largestPalindromic(num string) string {
	m := map[rune]int{}
	for _, n := range num {
		m[n] += 1
	}

	left := ""
	right := ""
	center := ""
	for _, k := range "9876543210" {
		v := m[k]
		for i := 0; i < v/2; i++ {
			left += string(k)
			right = string(k) + right
		}
		if v%2 == 1 {
			if center == "" {
				center = string(k)
			}
		}
	}

	if len(left) > 0 && left[0] == '0' {
		if center == "" {
			return "0"
		}
		return center
	}

	return left + center + right
}
