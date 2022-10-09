package leetcode

func hardestWorker(n int, logs [][]int) int {
	ct := 0
	max := 0
	ans := 0
	for _, v := range logs {
		id := v[0]
		t := v[1]
		if max < t-ct {
			max = t - ct
			ans = id
		}
		ct = t
	}

	return ans
}
