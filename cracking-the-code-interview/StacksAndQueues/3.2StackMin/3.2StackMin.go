package main

type node struct {
	value interface{}
	next  *node
	min   int
}

type queuable interface {
	getPriority() int
	//getNext() *queuable
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

func (n node) getNext() queuable {
	return queuable(n.next)
}

func (n node) setNext(newNext queuable) {
	nextNode := newNext.(node)
	n.next = &nextNode
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
	if pq.length == 1 {
		pq.first = &element
		return
	}

	element.setNext(*pq.first)

	if element.getPriority() < (*pq.first).getMin() {
		element.setMin(element.getPriority())
	} else {
		element.setMin((*pq.first).getMin())
	}

	pq.first = &element

}

// IMPLEMENTED AS A STACK
func (pq *priorityQueue) pop() queuable {

	popped := pq.first

	if popped != nil {

		nodable := (*popped).(node)
		newFirst := nodable.getNext()
		pq.first = &newFirst

		if pq.length > 0 {
			pq.length--
		}

		return *popped
	}

	return *popped
}
