package leetcode

/*
 * @lc app=leetcode id=838 lang=golang
 *
 * [838] Push Dominoes
 */

// @lc code=start
func pushDominoes(dominoes string) string {
	ds := []byte(dominoes)
	r := -1
	for i, d := range ds {
		if d == 'L' {
			if r == -1 {
				for j := i - 1; j >= 0 && ds[j] == '.'; j-- {
					ds[j] = 'L'
				}
			} else {
				j, k := r+1, i-1
				for j < k {
					ds[j] = 'R'
					ds[k] = 'L'
					j++
					k--
				}
				r = -1
			}
		}
		if d == 'R' {
			if r != -1 {
				for j := r + 1; j < i; j++ {
					ds[j] = 'R'
				}
			}
			r = i
		}
	}

	if r != -1 {
		for j := r + 1; j < len(dominoes); j++ {
			ds[j] = 'R'
		}
	}

	return string(ds)
}

// @lc code=end
