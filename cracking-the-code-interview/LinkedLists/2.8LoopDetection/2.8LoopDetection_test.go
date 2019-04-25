package main

import "testing"

func TestLoopDetection(t *testing.T) {
	node1 := LinkedListNode{value: 1}
	node2 := LinkedListNode{value: 2}
	node3 := LinkedListNode{value: 3}
	node4 := LinkedListNode{value: 4}
	node5 := LinkedListNode{value: 5}

	node1.next = &node2
	node2.next = &node3
	node3.next = &node4
	node4.next = &node5
	node5.next = &node2

	expected := true
	response := isLoop(&node1)

	if response != expected {
		t.Errorf("expected %t - got %t for loop in %+v", expected, response, node1)
	}

	// BREAK THE LOOP
	node5.next = nil

	newExpected := false
	newResponse := isLoop(&node1)

	if newResponse != newExpected {
		t.Errorf("expected %t - got %t for loop in %+v", newExpected, newResponse, node1)
		iterHead := &node1
		for iterHead != nil {
			t.Logf("%+v", iterHead)
			iterHead = iterHead.next
		}
	}

}
