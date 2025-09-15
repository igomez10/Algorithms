package tree

import "testing"

func TestTraverseDFS(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}

	initialTree := FromArray(input)

	traversedNodes := TraverseDFS(initialTree)

	traversedValues := []interface{}{}

	for i := range traversedNodes {
		traversedValues = append(traversedValues, traversedNodes[i].Val)
	}

	expectedOutput := []int{1, 2, 4, 5, 3, 6, 7}
	for i := range expectedOutput {

		if expectedOutput[i] != traversedValues[i] {

			baseMessage := "Unexpected value %d, expecting %d"
			t.Errorf(baseMessage, traversedValues[i], expectedOutput[i])
			t.Logf("%+v  != %+v", expectedOutput, traversedValues)
		}
	}

}
