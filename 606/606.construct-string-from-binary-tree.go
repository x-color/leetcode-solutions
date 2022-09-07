package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=606 lang=golang
 *
 * [606] Construct String from Binary Tree
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
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	l := tree2str(root.Left)
	r := tree2str(root.Right)
	if r != "" {
		return fmt.Sprintf("%v(%s)(%s)", root.Val, l, r)
	}
	if l != "" {
		return fmt.Sprintf("%v(%s)", root.Val, l)
	}
	return fmt.Sprint(root.Val)
}

// @lc code=end
