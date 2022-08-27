package leetcode

func isBadVersion(version int) bool {
	return true
}

/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		m := (r + l) / 2
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// @lc code=end
