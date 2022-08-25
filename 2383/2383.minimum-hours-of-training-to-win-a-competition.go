package leetcode

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	t := 0
	en := 0
	for _, n := range energy {
		en += n
	}
	if initialEnergy <= en {
		t = en - initialEnergy + 1
	}

	ex := experience[len(experience)-1] + 1
	for i := len(experience) - 2; i >= 0; i-- {
		cur := experience[i] + 1
		ex -= experience[i]
		if ex < cur {
			ex = cur
		}
	}

	if ex > initialExperience {
		t += ex - initialExperience
	}

	return t
}

// 4
// 47
