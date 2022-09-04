package leetcode

func checkDistances(s string, distance []int) bool {
	for i, r := range s {
		if d := distance[int(r-'a')]; d != -1 {
			ni := i + d + 1
			if ni >= len(s) || byte(r) != s[ni] {
				return false
			}
			distance[int(r-'a')] = -1
		}
	}

	return true
}
