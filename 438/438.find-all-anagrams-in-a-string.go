package leetcode

/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	cp := map[rune]int{}
	for _, r := range p {
		cp[r]++
	}

	ans := []int{}
	cs := map[rune]int{}
	for i, r := range s {
		cs[r]++
		if i-len(p) >= 0 {
			k := rune(s[i-len(p)])
			cs[k]--
			if cs[k] == 0 {
				delete(cs, k)
			}
		}

		if len(cs) != len(cp) {
			continue
		}
		same := true
		for k, v := range cp {
			if v != cs[k] {
				same = false
				break
			}
		}

		if same {
			ans = append(ans, i-len(p)+1)
		}

	}

	return ans
}

// @lc code=end
