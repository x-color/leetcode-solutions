package leetcode

/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
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
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddHead := head.Next
	even, odd := head, head.Next
	for odd != nil && odd.Next != nil {
		even.Next = odd.Next
		even = odd.Next
		odd.Next = even.Next
		odd = even.Next
	}

	even.Next = oddHead
	return head
}

// @lc code=end
