package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	return build(replace(bfs(root)))
}

func bfs(node *TreeNode) []int {
	left := 0
	nodes := []*TreeNode{node}
	for len(nodes[left:]) > 0 {
		l := len(nodes[left:])
		for i := 0; i < l; i++ {
			n := nodes[left+i]
			if n.Left != nil {
				nodes = append(nodes, n.Left, n.Right)
			}
		}
		left += l
	}

	rt := []int{}
	for _, v := range nodes {
		rt = append(rt, v.Val)
	}

	return rt
}

func build(nodes []int) *TreeNode {
	root := &TreeNode{
		Val: nodes[0],
	}
	nodes = nodes[1:]
	treeNodes := []*TreeNode{root}
	for len(nodes) > 0 {
		i := 0
		l := len(treeNodes)
		for j := 0; j < l; j++ {
			treeNodes[j].Left = &TreeNode{
				Val: nodes[i],
			}
			treeNodes[j].Right = &TreeNode{
				Val: nodes[i+1],
			}
			i += 2
			treeNodes = append(treeNodes, treeNodes[j].Left, treeNodes[j].Right)
		}
		treeNodes = treeNodes[l:]
		nodes = nodes[l*2:]
	}

	return root
}

func replace(nodes []int) []int {
	d := 1
	l, r := 0, 0
	for int(1<<d) < len(nodes) {
		l += int(1 << (d - 1))
		r += int(1 << d)
		if d%2 == 0 {
			d++
			continue
		}
		for i := 0; i < (r-l+1)/2; i++ {
			tmp := nodes[l+i]
			nodes[l+i] = nodes[r-i]
			nodes[r-i] = tmp
		}

		d++
	}

	return nodes
}
