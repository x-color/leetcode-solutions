package leetcode

func removeStars(s string) string {
	ans := make([]rune, 0, len(s))
	for _, r := range s {
		if r == '*' {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, r)
		}
	}
	return string(ans)
}

// leet**cod*e
// 00011000100

// i := strings.Index(s, "*")
// 		if i < 0 {
// 			return s
// 		}
// 		s = s[:i] + s[i+1:]
// 		for j := i - 1; j >= 0; j-- {
// 			if s[j] != '*' {
// 				s = s[:j] + s[j+1:]
// 				break
// 			}
// 		}
