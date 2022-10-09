package leetcode

func robotWithString(s string) string {
	count := make(map[rune]int, 0)
	for _, r := range s {
		count[r]++
	}

	min := 'a'
	p := make([]rune, 0, len(s))
	t := make([]rune, 0, len(s))
	for _, r := range s {
		t = append(t, r)
		count[r]--
		for len(t) > 0 {
			cr := rune(t[len(t)-1])
			for min <= cr && count[min] == 0 {
				min++
			}
			if cr > min {
				break
			}
			p = append(p, cr)
			t = t[:len(t)-1]
		}
	}

	for i := len(t) - 1; i >= 0; i-- {
		p = append(p, t[i])
	}

	return string(p)
}
