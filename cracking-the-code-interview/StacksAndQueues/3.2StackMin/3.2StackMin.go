package main

type node struct {
	value interface{}
	next  *node
	min   int
}

type queuable interface {
	getPriority() int
	getNext() *queuable
	setNext(queuable)
	getMin() int
	setMin(int)
}

type priorityQueue struct {
	first  *queuable
	length int
}

func (n node) getPriority() int {
	return n.value.(int)
}

func (n node) getNext() *node {
	return n.next
}

func (n *node) setNext(newNext node) {
	n.next = &newNext
}

func (n node) getMin() int {
	return n.min
}

func (n node) setMin(newMin int) {
	n.min = newMin
}

func main() {

}

// IMPLEMENTED AS A STACK
func (pq *priorityQueue) push(element queuable) {
	pq.length++

	element.setNext(*pq.first)

	if element.getPriority() < (*pq.first).getMin() {
		element.setMin(element.getPriority())
	} else {
		element.setMin((*pq.first).getMin())
	}

	pq.first = &element

}

// IMPLEMENTED AS A STACK
func (pq *priorityQueue) pop() *queuable {

	popped := pq.first

	if popped != nil {
		newFirst := (*pq.first).getNext()
		pq.first = newFirst
	}

	if pq.length > 0 {
		pq.length--
	}

	return popped
}
