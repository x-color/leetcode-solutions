package main

import "fmt"

// Word definitions
// - not increasing: not(all(x1 < x2 => f(x1) < f(x2)))
// - non-increasing: not(exist(x1 < x2 => f(x1) < f(x2))))
//
// See https://math.stackexchange.com/questions/115912/why-do-we-use-non-increasing-instead-of-decreasing

func goodIndices(nums []int, k int) []int {
	ni := make([]int, len(nums))
	for i := 1; i < len(nums)-k; i++ {
		if nums[i-1] >= nums[i] {
			ni[i] = ni[i-1] + 1
		}
	}
	nd := make([]int, len(nums))
	for i := len(nums) - 2; i > k; i-- {
		if nums[i] <= nums[i+1] {
			nd[i] = nd[i+1] + 1
		}
	}

	ans := []int{}
	for i := k; i < len(nums)-k; i++ {
		if ni[i-1]+1 >= k && nd[i+1]+1 >= k {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	fmt.Println("ANS:", goodIndices([]int{2, 1, 1, 1, 3, 4, 1}, 2), []int{2, 3})
	fmt.Println("ANS:", goodIndices([]int{2, 1, 1, 2}, 2), nil)
	fmt.Println("ANS:", goodIndices([]int{440043, 276285, 336957}, 1), []int{1})
	fmt.Println("ANS:", goodIndices([]int{388589, 17165, 726687, 401298, 600033, 537254, 301052, 151069, 399955}, 4), nil)
	fmt.Println("ANS:", goodIndices([]int{878724, 201541, 179099, 98437, 35765, 327555, 475851, 598885, 849470, 943442}, 4), []int{4, 5})
}
