package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	rs := make([]int, len(num1)+len(num2))
	for i := len(num2) - 1; i >= 0; i-- {
		n := int(num2[i] - '0')
		for j := len(num1) - 1; j >= 0; j-- {
			m := int(num1[j] - '0')
			rs[i+j+1] += n * m % 10
			rs[i+j] += n * m / 10
		}
	}

	ans := ""
	over := 0
	for i := len(rs) - 1; i >= 0; i-- {
		n := rs[i] + over
		ans = fmt.Sprintf("%d%s", n%10, ans)
		over = n / 10
	}

	for len(ans) > 0 && ans[0] == '0' {
		ans = ans[1:]
	}

	if ans == "" {
		return "0"
	}

	return ans
}

// @lc code=end
