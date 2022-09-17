package main

import "fmt"

/*
 * @lc app=leetcode id=336 lang=golang
 *
 * [336] Palindrome Pairs
 */

// This solution uses TrieTree.
// See
// - https://leetcode.com/problems/palindrome-pairs/discuss/79195/O(n-*-k2)-java-solution-with-Trie-structure
// - http://www.allenlipeng47.com/blog/index.php/2016/03/15/palindrome-pairs/

// @lc code=start
type TrieNode struct {
	next        map[byte]*TrieNode
	wordIdx     int
	palindromes []int
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		next:        make(map[byte]*TrieNode),
		wordIdx:     -1,
		palindromes: make([]int, 0),
	}
}

func (n *TrieNode) add(word string, idx int) {
	for i := len(word) - 1; i >= 0; i-- {
		if n.isPalindrome(word[:i+1]) {
			n.palindromes = append(n.palindromes, idx)
		}
		if n.next[word[i]] == nil {
			n.next[word[i]] = newTrieNode()
		}
		n = n.next[word[i]]
	}

	n.wordIdx = idx
}

func (n *TrieNode) search(word string, idx int) []int {
	rt := make([]int, 0)
	for i := range word {
		if n.wordIdx >= 0 && n.wordIdx != idx && n.isPalindrome(word[i:]) {
			rt = append(rt, n.wordIdx)
		}
		n = n.next[word[i]]
		if n == nil {
			return rt
		}
	}

	if n.wordIdx != -1 && n.wordIdx != idx {
		rt = append(rt, n.wordIdx)
	}

	for _, i := range n.palindromes {
		if i != idx {
			rt = append(rt, i)
		}
	}

	return rt
}

func (n TrieNode) isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func palindromePairs(words []string) [][]int {
	root := newTrieNode()

	for i, word := range words {
		root.add(word, i)
	}

	ans := make([][]int, 0)
	for i, word := range words {
		for _, j := range root.search(word, i) {
			ans = append(ans, []int{i, j})
		}
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(palindromePairs([]string{"bat", "tab", "cat"}))
	fmt.Println(palindromePairs([]string{"a", ""}))
	fmt.Println(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}))
}
