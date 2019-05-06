package tree

// TraverseDFS returns the node.Val from all the nodes in the tree by
// doing depth first search
func TraverseDFS(root *Node) []*Node {
	values := []*Node{root}

	if root.Left != nil {
		values = append(values, TraverseDFS(root.Left)...)
	}

	if root.Right != nil {
		values = append(values, TraverseDFS(root.Right)...)
	}

	return values

}
