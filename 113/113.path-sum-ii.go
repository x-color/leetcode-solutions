package leetcode

/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	return dfs(root, targetSum, make([]int, 0))
}

func dfs(node *TreeNode, n int, path []int) [][]int {
	if node == nil {
		return nil
	}

	n -= node.Val
	p := append(path, node.Val)
	if n == 0 && node.Left == nil && node.Right == nil {
		// append() returns the same pointer if the slice has a capacity
		// that it is possible to put a new value in the slice.
		// So, it must copy the slice to a new one to put it into ans.
		// If you don't copy it, dfs() rewrite the ans by the subsequent processing.
		r := make([]int, len(p))
		copy(r, p)
		return [][]int{r}
	}

	pl := dfs(node.Left, n, p)
	pr := dfs(node.Right, n, p)
	return append(pl, pr...)
}

// @lc code=end
