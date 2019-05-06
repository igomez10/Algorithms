package tree

import "testing"

func TestBFSTraversal(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7}

	tree := FromArray(values)

	traversedNodes := TraverseBFS(tree)

	traversedValues := []interface{}{}
	for i := range traversedNodes {
		traversedValues = append(traversedValues, traversedNodes[i].Val)
	}

	for i := range values {

		if values[i] != traversedValues[i] {
			baseMessage := "Expected %d got %d when traversing BFS"

			t.Errorf(baseMessage, values[i], traversedValues[i])

			t.Logf("%+v    !=    %+v ", values, traversedValues)
		}

	}
}
