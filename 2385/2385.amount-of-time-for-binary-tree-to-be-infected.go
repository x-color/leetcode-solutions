package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val   int
	Nodes []*Node
}

func amountOfTime(root *TreeNode, start int) int {
	_, rn := graph(nil, root, start)
	return depth(nil, rn, 0)
}

func depth(p, cur *Node, d int) int {
	max := d
	for _, n := range cur.Nodes {
		if n == p || n == nil {
			continue
		}
		dep := depth(cur, n, d+1)
		if dep > max {
			max = dep
		}
	}
	return max
}

func graph(p *Node, n *TreeNode, v int) (*Node, *Node) {
	if n == nil {
		return nil, nil
	}
	node := &Node{
		Val:   n.Val,
		Nodes: []*Node{p},
	}

	var root *Node
	if l, rn := graph(node, n.Left, v); l != nil {
		node.Nodes = append(node.Nodes, l)
		if rn != nil {
			root = rn
		}
	}
	if r, rn := graph(node, n.Right, v); r != nil {
		node.Nodes = append(node.Nodes, r)
		if rn != nil {
			root = rn
		}
	}

	if v == node.Val {
		root = node
	}

	return node, root
}
