package leetcode

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	visit := make([][]bool, len(grid))
	for i := range visit {
		visit[i] = make([]bool, len(grid[0]))
	}

	c := 0
	for y := range grid {
		for x := range grid[y] {
			if dfs(x, y, grid) {
				c++
			}
		}
	}

	return c
}

func dfs(x, y int, grid [][]byte) bool {
	if grid[y][x] != '1' {
		return false
	}
	grid[y][x] = '2'

	if x-1 >= 0 {
		dfs(x-1, y, grid)
	}
	if x+1 < len(grid[0]) {
		dfs(x+1, y, grid)
	}
	if y-1 >= 0 {
		dfs(x, y-1, grid)
	}
	if y+1 < len(grid) {
		dfs(x, y+1, grid)
	}

	return true
}

// @lc code=end
