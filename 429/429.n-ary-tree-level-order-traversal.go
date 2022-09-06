package leetcode

type Node struct {
	Val      int
	Children []*Node
}

/*
 * @lc app=leetcode id=429 lang=golang
 *
 * [429] N-ary Tree Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	nodes := []*Node{root}
	level := 0
	for len(nodes) > 0 {
		ans = append(ans, []int{})
		l := len(nodes)
		for j := 0; j < l; j++ {
			ans[level] = append(ans[level], nodes[j].Val)
			for _, node := range nodes[j].Children {
				nodes = append(nodes, node)
			}
		}
		nodes = nodes[l:]
		level++
	}

	return ans
}

// @lc code=end
