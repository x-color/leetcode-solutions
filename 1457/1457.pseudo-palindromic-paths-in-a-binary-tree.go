package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=1457 lang=golang
 *
 * [1457] Pseudo-Palindromic Paths in a Binary Tree
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
func pseudoPalindromicPaths(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, path int) int {
	if node == nil {
		return 0
	}
	counter := 0
	path = path ^ (1 << node.Val)

	counter += dfs(node.Left, path)
	counter += dfs(node.Right, path)
	if node.Left == nil && node.Right == nil {
		// Check whether path is pseudo-palindromic path
		// This condition is passed if path has only one 1 bit.
		// path = 0100: path & (path - 1) = 0100 & 0011 = 0
		// path = 0110: path & (path - 1) = 0100 & 0101 = 0100 (not 0)
		if path&(path-1) == 0 {
			counter++
		}
	}
	return counter
}

// @lc code=end
