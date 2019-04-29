package main

import "testing"

func TestPriorityQueue(t *testing.T) {
	pq := priorityQueue{}

	node1 := node{value: 1, min: 0}
	node2 := node{value: 2, min: 1}
	node3 := node{value: 3, min: 2}
	node4 := node{value: 4, min: 3}
	node5 := node{value: 5, min: 4}
	node6 := node{value: 6, min: 5}

	node1.next = &node2
	node2.next = &node3
	node3.next = &node4
	node4.next = &node5
	node5.next = &node6

	pq.push(node1)
	if (*pq.first).getPriority() != node1.getPriority() {
		t.Errorf("Expected first element in queue to be %d - got %d", node1.value, (*pq.first).getPriority())
	}

	expectedLength := 1
	if pq.length != 1 {
		t.Errorf("Expected length of queue %+d - got %d", expectedLength, pq.length)
	}

	pq.push(node2)
	if (*pq.first).getPriority() != node2.value {
		t.Errorf("Expected first element in queue to be %d - got %d", node2.value, (*pq.first).getPriority())
	}

	popped := pq.pop()
	if popped.getPriority() != node2.value {
		t.Errorf("expected to pop %+v - got %+v", node2.value, popped.getPriority())
	}

	secondPopped := pq.pop()
	if secondPopped.getPriority() != node1.value {
		t.Errorf("expected to pop %+v - got %+v", node1.value, popped.getPriority())
	}

	nilPopped := pq.pop()
	if nilPopped != nil {
		t.Errorf("expected to pop %+v - got %+v", node1.value, popped.getPriority())
	}
}

//func TestNodeSetNext(t *testing.T) {
//
//	node1 := node{value: 1, min: 0}
//	node2 := node{value: 2, min: 1}
//
//	node1.setNext(node2)
//
//	if node1.getNext() != node2 {
//		t.Errorf("Error assigning next node to node, expected %+v got %+v", node2, node1.getNext())
//	}
//}
