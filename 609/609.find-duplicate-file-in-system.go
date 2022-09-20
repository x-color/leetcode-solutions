package leetcode

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=609 lang=golang
 *
 * [609] Find Duplicate File in System
 */

// @lc code=start
func findDuplicate(paths []string) [][]string {
	files := map[string][]string{}
	for _, p := range paths {
		s := strings.Split(p, " ")
		dir := s[0]
		for _, f := range s[1:] {
			fc := strings.Split(f, "(")
			file := fc[0]
			content := fc[1][:len(fc[1])-1]
			if _, ok := files[content]; !ok {
				files[content] = make([]string, 0)
			}
			files[content] = append(files[content], fmt.Sprintf("%s/%s", dir, file))
		}
	}

	ans := [][]string{}
	for _, fs := range files {
		if len(fs) < 2 {
			continue
		}
		ans = append(ans, fs)
	}

	return ans
}

// @lc code=end
