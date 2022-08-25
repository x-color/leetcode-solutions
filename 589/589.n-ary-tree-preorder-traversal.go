package leetcode

type Node struct {
	Val      int
	Children []*Node
}

/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	r := []int{}
	if root == nil {
		return r
	}

	r = append(r, root.Val)
	for _, n := range root.Children {
		r = append(r, preorder(n)...)
	}
	return r
}

// @lc code=end
