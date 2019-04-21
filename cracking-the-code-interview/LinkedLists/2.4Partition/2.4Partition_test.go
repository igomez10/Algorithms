package main

import "strings"
import "testing"

func TestLLFromArray(t *testing.T) {
	values1 := []int{1, 1, 2, 2, 3, 4, 5, 6, 6, 7, 7, 2, 1, 2, 3, 4, 5, 6, 7}

	head := LLFromArray(values1)

	for i := 0; i < len(values1); i++ {
		if values1[i] != head.value {
			t.Errorf("Expected %d got %d", values1[i], head.value)
		}
		head = head.next
	}
}

func TestLinkedListNode_ToString(t *testing.T) {
	values := []int{1, 1, 2, 2, 3, 3, 4, 6, 5, 4}
	head := LLFromArray(values)
	stringValue := head.ToString()
	expectedValue := "1122334654"
	if strings.Compare(stringValue, expectedValue) != 0 {
		t.Errorf("expected %s got %s", stringValue, expectedValue)
	}
}

func TestPartition(t *testing.T) {

	values := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	partitionValue := 5

	expectedValues := []int{1, 2, 3, 4, 5, 9, 8, 7, 6}

	head := LLFromArray(values)

	partition(&head, partitionValue)

	for i := 0; i < len(expectedValues); i++ {
		if expectedValues[i] != head.value {
			t.Errorf("Expected %d, got %d", expectedValues[i], head.value)
			t.Error(head.ToString())
		}
		head = head.next
	}

}
