package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	level := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		if level%2 == 1 {
			l, r := 0, len(nodes)-1
			for l < r {
				nodes[l].Val, nodes[r].Val = nodes[r].Val, nodes[l].Val
				l++
				r--
			}
		}

		length := len(nodes)
		for i := 0; i < length; i++ {
			if nodes[i].Left == nil {
				break
			}
			nodes = append(nodes, nodes[i].Left, nodes[i].Right)
		}
		nodes = nodes[length:]
		level++
	}
	return root
}
