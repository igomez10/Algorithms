package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// FromArray returns a binary complete tree from the given array
func FromArray(arr []int) *Node {
	root := InsertLevelOrder(nil, arr, 0, len(arr))
	return root
}

func (t *Node) getAllChildren() []*Node {
	response := []*Node{}

	pendingToVisit := []*Node{t}

	for len(pendingToVisit) > 0 {

		response = append(response, pendingToVisit[0])
		pendingToVisit = pendingToVisit[1:]

		if t.Left != nil {
			leftChildren := t.Left.getAllChildren()
			response = append(response, leftChildren...)
		}

		if t.Right != nil {
			rightChildren := t.Right.getAllChildren()
			response = append(response, rightChildren...)
		}

	}
	return response
}
