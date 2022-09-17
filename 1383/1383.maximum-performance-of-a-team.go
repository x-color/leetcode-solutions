package leetcode

import (
	"container/heap"
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=1383 lang=golang
 *
 * [1383] Maximum Performance of a Team
 */

// @lc code=start

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	engineers := make([][2]int, n)
	for i := 0; i < n; i++ {
		engineers[i] = [2]int{efficiency[i], speed[i]}
	}
	sort.Slice(engineers, func(i, j int) bool {
		return engineers[i][0] > engineers[j][0]
	})

	totalSpeed, res := 0, math.MinInt32
	pq := &IntHeap{}
	heap.Init(pq)
	for _, engineer := range engineers {
		heap.Push(pq, engineer[1])
		totalSpeed += engineer[1]
		if pq.Len() > k {
			totalSpeed -= heap.Pop(pq).(int)
		}
		if totalSpeed*engineer[0] > res {
			res = totalSpeed * engineer[0]
		}
	}
	return res % int(1e9+7)
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
