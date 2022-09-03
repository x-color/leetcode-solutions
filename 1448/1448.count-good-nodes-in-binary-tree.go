package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=1448 lang=golang
 *
 * [1448] Count Good Nodes in Binary Tree
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
func goodNodes(root *TreeNode) int {
	return dfs(root, -100000)
}

func dfs(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}

	if node.Val < max {
		return dfs(node.Right, max) + dfs(node.Left, max)
	}

	return dfs(node.Right, node.Val) + dfs(node.Left, node.Val) + 1
}

// @lc code=end
