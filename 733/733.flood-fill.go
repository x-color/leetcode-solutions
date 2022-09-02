package leetcode

/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 */

// @lc code=start
type pos struct {
	r int
	c int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	s := image[sr][sc]
	visited := make([][]int, len(image))
	for i := range visited {
		visited[i] = make([]int, len(image[i]))
	}
	h, w := len(image), len(image[0])
	q := []pos{
		{
			r: sr,
			c: sc,
		},
	}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		if image[c.r][c.c] != s || visited[c.r][c.c] == 1 {
			continue
		}
		image[c.r][c.c] = color
		visited[c.r][c.c] = 1
		q = append(q, neighbors(c, h, w)...)
	}
	return image
}

func neighbors(p pos, h, w int) []pos {
	ps := make([]pos, 0, 4)
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range dir {
		n := pos{
			r: p.r + d[0],
			c: p.c + d[1],
		}
		if n.r < 0 || n.r > h-1 {
			continue
		}
		if n.c < 0 || n.c > w-1 {
			continue
		}
		ps = append(ps, n)
	}
	return ps
}

// @lc code=end
