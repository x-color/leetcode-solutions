package leetcode

func garbageCollection(garbage []string, travel []int) int {
	t := map[string]int{}
	tmp := map[string]int{}
	for i := range garbage {
		for _, r := range garbage[i] {
			t[string(r)] += tmp[string(r)] + 1
			tmp[string(r)] = 0
		}

		if i < len(travel) {
			tmp["M"] += travel[i]
			tmp["P"] += travel[i]
			tmp["G"] += travel[i]
		}
	}

	s := 0
	for _, v := range t {
		s += v
	}
	return s
}

// G: 1, 0, 1, 2
// P: 0, 1, 1, 0
