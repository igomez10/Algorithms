package tree

// TraverseBFS returns the node.Val from all the nodes in the tree by
// doing in level search
func TraverseBFS(root *Node) []interface{} {
	values := []interface{}{root.Val}

	if root.Left != nil {
		values = append(values, TraverseBFS(root.Left)...)
	}

	if root.Right != nil {
		values = append(values, TraverseBFS(root.Right)...)
	}

	return values
}

// InsertLevelOrder will insert a new node in such a way that if the initial given
// tree was balanced, the returned tree is also balanced.
func InsertLevelOrder(root *Node, arr []int, i, n int) *Node {
	if i < n {
		temp := Node{Val: arr[i]}
		root = &temp

		// insert left child
		root.Left = InsertLevelOrder(root.Left, arr, 2*i+1, n)

		// insert right child
		root.Right = InsertLevelOrder(root.Right, arr, 2*i+2, n)
	}
	return root
}
