package tree

import (
	"fmt"
	"strings"
)

// Node represents a node in the tree, has a value int and might have left
// and right child
type Node struct {
	Val   interface{}
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

// GetHeight returns the longest path from root to leaf, counting root as level 1
func GetHeight(root Node) int {
	counter := 0
	for iterativeRoot := &root; iterativeRoot != nil; iterativeRoot = iterativeRoot.Left {
		counter++
	}
	return counter
}

// toString returns a string version of the given tree,
// intentend for debugging and human readable format
func toString(root *Node) string {
	builder := strings.Builder{}

	traversedNodes := TraverseBFS(root)

	nodesInCurrentLevel := 1

	for len(traversedNodes) > 0 {

		for i := range traversedNodes[:nodesInCurrentLevel] {
			currentVal := fmt.Sprintf("%+v", traversedNodes[i].Val)
			builder.WriteString(currentVal)
		}
		traversedNodes = traversedNodes[nodesInCurrentLevel:]
		nodesInCurrentLevel *= 2
		builder.WriteRune('\n')
	}
	return builder.String()
}
