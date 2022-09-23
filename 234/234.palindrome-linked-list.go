package leetcode

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
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
func isPalindrome(head *ListNode) bool {
	var reverse *ListNode
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		tmp := reverse
		reverse = slow
		slow = slow.Next
		reverse.Next = tmp
	}

	if fast != nil {
		slow = slow.Next
	}

	for reverse != nil && reverse.Val == slow.Val {
		reverse = reverse.Next
		slow = slow.Next
	}

	return reverse == nil
}

// @lc code=end
