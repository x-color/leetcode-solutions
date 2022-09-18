package leetcode

/*
 * @lc app=leetcode id=1706 lang=golang
 *
 * [1706] Where Will the Ball Fall
 */

// @lc code=start
func findBall(grid [][]int) []int {
	balls := make([]int, len(grid[0]), len(grid[0]))
	for i := range balls {
		balls[i] = i
	}

	for _, row := range grid {
		for i, p := range balls {
			if p == -1 {
				continue
			}
			switch {
			case row[p] == 1 && right(p, row):
				balls[i]++
			case row[p] == -1 && left(p, row):
				balls[i]--
			default:
				balls[i] = -1
			}
		}
	}
	return balls
}

func right(p int, row []int) bool {
	if p+1 < len(row) {
		return row[p+1] == 1
	}

	return false
}

func left(p int, row []int) bool {
	if p-1 >= 0 {
		return row[p-1] == -1
	}

	return false
}

// @lc code=end
