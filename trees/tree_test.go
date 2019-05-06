package tree

import "testing"

func TestTreeFromArray(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7}

	tree := FromArray(values)

	arrayValues := tree.getAllChildren()

	inputFrequencyTable := make(map[interface{}]int)
	for i := range arrayValues {
		inputFrequencyTable[values[i]]++
	}

	// Compare if elements given in the initial input are in the new tree, dont
	// worry about order, only if present with the same frequency
	frequencyTable := make(map[interface{}]int)
	for i := range arrayValues {
		frequencyTable[(*arrayValues[i]).Val]++
	}

	for inputKey, inputValue := range inputFrequencyTable {
		if valueInTree, exist := frequencyTable[inputKey]; exist {
			if valueInTree != inputValue {
				t.Errorf("Frenquency of value %d expected to be %d but got %d", inputKey, inputValue, valueInTree)
			}
		} else {
			t.Errorf("Value %+d not found in tree", inputKey)
		}
	}

	// TRIVIAL TESTS

	if tree.Val != 1 {
		t.Errorf("Unepected root val %+v, expected %d", tree.Val, 1)
	}

	if tree.Left.Left.Val != 4 {
		t.Errorf("Unepected root val %+v, expected %d", tree.Val, 4)
	}

	if tree.Right.Left.Val != 6 {
		t.Errorf("Unepected root val %+v, expected %d", tree.Val, 6)
	}

}
