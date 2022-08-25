package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			slow = head
			for {
				if fast == slow {
					return slow
				}
				slow = slow.Next
				fast = fast.Next
			}
		}
	}
	return nil
}

// @lc code=end

// slow := head
// fast := head
// for fast != nil && fast.Next != nil {
// 	slow = slow.Next
// 	fast = fast.Next.Next
// 	if slow == fast {
// 		return fast
// 	}
// }
