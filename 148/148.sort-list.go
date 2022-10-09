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

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return mergeSort(head)
}

func split(head *ListNode) (*ListNode, *ListNode) {
	var tmp *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		tmp = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp.Next = nil

	return head, slow
}

func mergeSort(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	a, b := split(head)
	a = mergeSort(a)
	b = mergeSort(b)
	root := &ListNode{
		Next: nil,
	}
	cur := root
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}

	if a != nil {
		cur.Next = a
	} else {
		cur.Next = b
	}

	return root.Next
}

// @lc code=end

// This is an invalid answer
// It causes Time Limit Exceeded (Time=O(n^2))
// func sortList(head *ListNode) *ListNode {
// 	root := &ListNode{Next: head}
// 	for n := root; n.Next != nil; n = n.Next {
// 		for m := n.Next; m.Next != nil; m = m.Next {
// 			// Swap Nodes (not only nodes' value)
// 			// n.Next <-> m.Next
// 			if n.Next.Val > m.Next.Val {
// 				nn := n.Next
// 				mnn := m.Next.Next
// 				if nn.Next == m.Next {
// 					n.Next = m.Next
// 					m.Next = m.Next.Next
// 					n.Next.Next = nn
// 					m = n.Next
// 				} else {
// 					n.Next = m.Next
// 					m.Next.Next = nn.Next
// 					m.Next = nn
// 					nn.Next = mnn
// 				}
// 			}
// 		}
// 	}

// 	return root.Next
// }
