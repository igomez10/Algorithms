package main

import "testing"

func TestIntersection(t *testing.T) {
	node1 := LinkedListNode{value: 1}
	node2 := LinkedListNode{value: 2}
	node3 := LinkedListNode{value: 3}
	node4 := LinkedListNode{value: 4}

	node1.next = &node3
	node2.next = &node3
	node3.next = &node4

	intersected := Intersection(&node1, &node2)
	if intersected != &node3 {
		t.Errorf("expected node node3 - got %+v", intersected)
	}

	nodeA := LinkedListNode{value: 1}
	nodeB := LinkedListNode{value: 2}
	nodeC := LinkedListNode{value: 3}
	nodeD := LinkedListNode{value: 4}
	nodeE := LinkedListNode{value: 5}

	nodeA.next = &nodeB
	nodeB.next = &nodeC
	nodeC.next = &nodeD
	nodeD.next = &nodeE

	nonIntersected := Intersection(&nodeA, &node1)
	if nonIntersected != nil {
		t.Errorf("expected nil intersection, got %+v", nonIntersected)
	}
}
