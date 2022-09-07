package leetcode

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=987 lang=golang
 *
 * [987] Vertical Order Traversal of a Binary Tree
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
// Triplet for each element
type Node struct {
	col int
	row int
	val int
}

func verticalTraversal(root *TreeNode) [][]int {
	nodes := dfs(root, 0, 0, []Node{})

	// Less function for lexicograplical order
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].col != nodes[j].col {
			return nodes[i].col < nodes[j].col
		} else if nodes[i].row != nodes[j].row {
			return nodes[i].row < nodes[j].row
		} else if nodes[i].val != nodes[j].val {
			return nodes[i].val < nodes[j].val
		} else {
			return false
		}
	})

	// Grouping sorted elements
	ans := make([][]int, 0)
	prev := nodes[0].col + 1 // hack to handle first element
	for _, node := range nodes {
		if node.col != prev {
			ans = append(ans, []int{})
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], node.val)

		prev = node.col
	}

	return ans
}

// Simple DFS
func dfs(root *TreeNode, row int, col int, nodes []Node) []Node {
	if root == nil {
		return nodes
	}

	nodes = append(nodes, Node{col, row, root.Val})
	nodes = dfs(root.Left, row+1, col-1, nodes)
	nodes = dfs(root.Right, row+1, col+1, nodes)

	return nodes
}

// @lc code=end
