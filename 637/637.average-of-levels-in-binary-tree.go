package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	ans := []float64{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		l := len(nodes)
		s := 0
		c := 0.0
		for i := 0; i < l; i++ {
			s += nodes[i].Val
			c++
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}
		ans = append(ans, float64(s)/c)
		nodes = nodes[l:]
	}

	return ans
}

// @lc code=end
