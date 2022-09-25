package leetcode

import "sort"

type P struct {
	name   string
	height int
}

func sortPeople(names []string, heights []int) []string {
	ps := []P{}
	for i := range names {
		ps = append(ps, P{names[i], heights[i]})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].height > ps[j].height
	})

	ans := []string{}
	for _, p := range ps {
		ans = append(ans, p.name)
	}
	return ans
}
