package leetcode

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	rt := make([]int, 0, len(matrix)*len(matrix[0]))

	x, y := 0, 0
	rt = append(rt, matrix[y][x])
	matrix[y][x] = -101
	for {
		sx, sy := x, y
		for x+1 < len(matrix[0]) && matrix[y][x+1] > -101 {
			x++
			rt = append(rt, matrix[y][x])
			matrix[y][x] = -101
		}
		for y+1 < len(matrix) && matrix[y+1][x] > -101 {
			y++
			rt = append(rt, matrix[y][x])
			matrix[y][x] = -101
		}
		for x-1 >= 0 && matrix[y][x-1] > -101 {
			x--
			rt = append(rt, matrix[y][x])
			matrix[y][x] = -101
		}
		for y-1 >= 0 && matrix[y-1][x] > -101 {
			y--
			rt = append(rt, matrix[y][x])
			matrix[y][x] = -101
		}
		if x == sx && y == sy {
			return rt
		}
	}
}

// @lc code=end
