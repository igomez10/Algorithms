package tree

// TraverseBFS returns the node.Val from all the nodes in the tree by
// doing in level search
func TraverseBFS(root *Node) []*Node {
	values := []*Node{}

	valuesToVisit := []*Node{root}

	for len(valuesToVisit) > 0 {

		values = append(values, valuesToVisit[0])

		if valuesToVisit[0].Left != nil {
			valuesToVisit = append(valuesToVisit, valuesToVisit[0].Left)
		}
		if valuesToVisit[0].Right != nil {
			valuesToVisit = append(valuesToVisit, valuesToVisit[0].Right)
		}
		valuesToVisit = valuesToVisit[1:] //dequeue
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
