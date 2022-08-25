package leetcode

/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := s2t[s[i]]; !ok {
			s2t[s[i]] = t[i]
		}
		if _, ok := t2s[t[i]]; !ok {
			t2s[t[i]] = s[i]
		}

		if s2t[s[i]] != t[i] || t2s[t[i]] != s[i] {
			return false
		}
	}
	return true
}

// @lc code=end
