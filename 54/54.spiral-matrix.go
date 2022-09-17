package leetcode

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	rt := make([]int, 0, len(matrix)*len(matrix[0]))
	x, y, dict := 0, 0, 1
	for {
		rt = append(rt, matrix[y][x])
		matrix[y][x] = -101
		x, y, dict = next(x, y, dict, matrix)
		if dict == 0 {
			return rt
		}
	}
}

func next(x, y, dict int, matrix [][]int) (int, int, int) {
	switch dict {
	case 1:
		if x+1 < len(matrix[0]) && matrix[y][x+1] > -101 {
			return x + 1, y, 1
		}
		fallthrough
	case 2:
		if y+1 < len(matrix) && matrix[y+1][x] > -101 {
			return x, y + 1, 2
		}
		fallthrough
	case 3:
		if x-1 >= 0 && matrix[y][x-1] > -101 {
			return x - 1, y, 3
		}
		fallthrough
	case 4:
		if y-1 >= 0 && matrix[y-1][x] > -101 {
			return x, y - 1, 4
		}
		fallthrough
	default:
		if x+1 < len(matrix[0]) && matrix[y][x+1] > -101 {
			return x + 1, y, 1
		}
	}

	return x, y, 0
}

// @lc code=end
