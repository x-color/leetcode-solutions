package leetcode

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// This is an invalid answer
// It causes Time Limit Exceeded (Time=O(n^2))
func sortList(head *ListNode) *ListNode {
	root := &ListNode{Next: head}
	for n := root; n.Next != nil; n = n.Next {
		for m := n.Next; m.Next != nil; m = m.Next {
			// Swap Nodes (not only nodes' value)
			// n.Next <-> m.Next
			if n.Next.Val > m.Next.Val {
				nn := n.Next
				mnn := m.Next.Next
				if nn.Next == m.Next {
					n.Next = m.Next
					m.Next = m.Next.Next
					n.Next.Next = nn
					m = n.Next
				} else {
					n.Next = m.Next
					m.Next.Next = nn.Next
					m.Next = nn
					nn.Next = mnn
				}
			}
		}
	}

	return root.Next
}

// @lc code=end
