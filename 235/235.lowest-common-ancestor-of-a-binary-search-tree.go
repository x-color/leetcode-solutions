package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if q.Val < p.Val {
		p, q = q, p
	}
	if root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}

// @lc code=end
