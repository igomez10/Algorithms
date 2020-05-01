package main

import (
	"fmt"
	"log"
)

type linkedList interface {
	printList()
	getNext() *linkedList
	getData() interface{}
}

type doublyLinkedListNode struct {
	data int32
	next *doublyLinkedListNode
	prev *doublyLinkedListNode
}

type singlyLinkedListNode struct {
	data int32
	next *singlyLinkedListNode
}

func (list singlyLinkedListNode) getData() interface{} {
	return list.data
}

func (list doublyLinkedListNode) getData() interface{} {
	return list.data
}

func (list *singlyLinkedListNode) getNext() *linkedList {
	ans := linkedList(list.next)
	return &ans
}

func (list *doublyLinkedListNode) getNext() *linkedList {
	ans := linkedList(list.next)
	return &ans
}

func createDoublyLLFromSinlgyLL(head *singlyLinkedListNode) *doublyLinkedListNode {
	newHead := &doublyLinkedListNode{
		data: head.data,
	}
	lastElement := newHead
	for head != nil {

		lastElement.next = &doublyLinkedListNode{
			data: head.data,
			prev: lastElement,
		}

		lastElement = lastElement.next
		head = head.next
	}
	return newHead
}

func main() {

	fmt.Print("\n\n\n\n\n")
	initalValues := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	firstNode := createSinglyLinkedList(initalValues)
	head := firstNode
	// COUNT
	fmt.Println(count(head))
	fmt.Printf("The list is made of %d elements -- %v \n", count(head), count(head) == int32(len(initalValues)))

	if count(head) != int32(len(initalValues)) {
		message := "%s is not working"
		panic(fmt.Sprintf(message, count))
	}
	//GET Nth
	nth := int32(5)
	nthNode := getNodeFromTail(head, nth)
	nthInArray := initalValues[int32(len(initalValues))-nth]
	fmt.Printf("The %dth element in the list is %d == %d -- %t \n", nth, nthNode, nthInArray, nthInArray == nthNode)

	//DELETE LIST
	//deleteList(&head)
	//printList(head)

	// POP
	//	pop(&head)
	//	stringMold := "Pop element from list, new length is now %d (%t), first element is now %d == %d -- %t \n"
	//	isWorking := head.data == initalValues[1]
	//	message := fmt.Sprintf(stringMold, count(head), count(head) == int32(len(initalValues)-1), head.data, initalValues[1], isWorking)
	//	if !isWorking {
	//      printList(head)
	//		panic(message) //	} else {
	//		fmt.Println(message)
	//
	//	}

	// PUSH
	//	push(&head, initalValues[0])
	//	stringMold = "Push element %d to list, new length is now %d (%t), first element is now %d == %d -- %t \n"
	//	isWorking = head.data == initalValues[0]
	//	message = fmt.Sprintf(stringMold, initalValues[0], count(head), count(head) == int32(len(initalValues)), head.data, initalValues[0], isWorking)
	//	if !isWorking {
	//      printList(head)
	//		panic(message)
	//	} else {
	//		fmt.Println(message)
	//	}

	// INSERT IN NTH POSITION
	insertedElement := int32(10)
	index := 9
	insertNth(head, index, insertedElement)
	stringMold := "Inserted element %d to list in index %d, new length is now %d  -- Expected %d got %d -- %t\n"
	isWorking := getNode(head, int32(index)) == int32(insertedElement)
	message := fmt.Sprintf(stringMold, insertedElement, index, count(head), insertedElement, getNode(head, int32(index)), isWorking)
	if !isWorking {
		(*head).printList()
		panic(message)
	} else {
		fmt.Println(message)
	}

	// SORTED INSERT
	insertedElement = int32(11)
	sortedInsert(head, insertedElement)

	// APPEND
	firstArray := []int32{1, 2, 3, 4, 5, 6}
	secondArray := []int32{7, 8, 9, 10, 11, 12, 13}
	firstList := createSinglyLinkedList(firstArray)
	secondList := createSinglyLinkedList(secondArray)
	appendLinkedLists(firstList, secondList)
	receivedLength := count(firstList)
	expectedLength := int32(len(firstArray) + len(secondArray))
	moldMessage := "Length of linked list, expected %d and got %d -- %t\n"
	fmt.Printf(moldMessage, expectedLength, receivedLength, expectedLength == receivedLength)
	if count(firstList) != expectedLength {
		firstList.printList()
		panic("Error appending lists")
	}

	//FRONTBACKSPLIT
	initialArray := []int32{2, 3, 5, 7, 11}
	newHead := createSinglyLinkedList(initialArray)
	splittedHead := frontBackSplit(newHead)
	expectedLengthFirstList := int(len(initialArray)/2) + len(initialArray)%2
	expectedLengthSecondList := len(initialArray) - expectedLengthFirstList
	receivedLengthFirstList := count(newHead)
	receivedLengthSecondList := count(*splittedHead)
	moldMessage = "Splitted list first list has expected len %d and received %d , secondList expected len %d and received %d -- %t"
	isWorking = int32(expectedLengthFirstList) == receivedLengthFirstList && int32(expectedLengthSecondList) == receivedLengthSecondList
	message = fmt.Sprintf(moldMessage, expectedLengthFirstList, receivedLengthFirstList, expectedLengthSecondList, receivedLengthSecondList, isWorking)
	fmt.Println(message)

	if !isWorking {
		newHead.printList()
		fmt.Println()
		(*splittedHead).printList()
		panic(message)
	}

	//REMOVE DUPLICATES
	initialArray = []int32{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 6, 6, 7, 8, 9, 9}
	newHead = createSinglyLinkedList(initialArray)
	removeDuplicates(newHead)
	if !compareArrayWithLinkedList(buildSetFromArray(initialArray), newHead) {
		newHead.printList()
		fmt.Println(buildSetFromArray(initialArray))
		fmt.Printf("Length of set array is %d while length of set list %d", len(buildSetFromArray(initialArray)), count(newHead))
		panic("The set of array of linked list and array is not the same")

	}
	//MOVE NODE
	firstArray = []int32{1, 2, 3}
	secondArray = []int32{1, 2, 3}

	firstList = createSinglyLinkedList(firstArray)
	secondList = createSinglyLinkedList(secondArray)
	moveNode(firstList, secondList)
	if count(firstList) != count(secondList)+1 {
		panic("Unexpected move for node")
	}

	//ALTERNATING SPLIT
	initialArray = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newHead = createSinglyLinkedList(initialArray)
	generatedHead := alternatingSplit(&newHead)
	if count(newHead)+count(*generatedHead) != int32(len(initialArray)) {
		panic("Coult not alternate split")
	}

	firstArray = []int32{2, 4, 6, 7, 8, 9, 10}
	secondArray = []int32{1, 3, 5}
	firstList = createSinglyLinkedList(firstArray)
	secondList = createSinglyLinkedList(secondArray)
	shuffleMerge(firstList, &secondList)

	if count(firstList) != int32(len(firstArray)+len(secondArray)) && count(secondList) != 0 {
		panic("Error doing alternating split")
	}

	//SORTED MERGE

	firstArray = []int32{1, 5, 7, 8, 9, 10, 11, 12, 13}
	secondArray = []int32{2, 3, 4, 6}

	firstList = createSinglyLinkedList(firstArray)
	secondList = createSinglyLinkedList(secondArray)
	newList := sortedMerge(firstList, secondList)
	if int32(len(firstArray)+len(secondArray)) != count(*newList) {
		(*newList).printList()
		panic("Error merging list, dimensions dont match")
	}

	firstArray = []int32{6, 5, 4, 3, 1, 2, 1, 9, -5}
	firstList = createSinglyLinkedList(firstArray)
	mergeSort(firstList)
	currentHead := firstList
	for currentHead.getNext() != nil {
		currentHeadData := currentHead.getData().(int32)
		nextNodeData := (*currentHead.getNext()).getData().(int32)
		if !(currentHeadData <= nextNodeData) {
			message := fmt.Sprintf("ERROR SORTING LL, found %d before %d", currentHeadData, nextNodeData)
			panic(message)
		}
		currentHead = (*currentHead.getNext())
	}
	if len(firstArray) != int(count(firstList)) {
		message := fmt.Sprintf("Expected %d and received %d", len(firstArray), count(firstList))
		panic("Error in dimension of sorted list: \n\t " + message)
	}
	fmt.Println("Lenght of list worked -- TRUE")

	//	//SORTED INTERSECT
	//	firstArray = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//	secondArray = []int32{2, 4, 6, 8}
	//
	//	firstList = createSinglyLinkedList(firstArray)
	//	secondList = createSinglyLinkedList(secondArray)
	//	newHead = sortedIntersect(&firstList, &secondList)
	//	commonElementsInArrays := findCommonInArray(firstArray, secondArray)
	//	if len(commonElementsInArrays) != int(count(newHead)) {
	//		message := fmt.Sprintf("Error checking common elements in linked lists, expected %d and received %d", len(commonElementsInArrays), count(newHead))
	//		panic(message)
	//	}
	//
	//	// REVERSE
	//	firstArray = []int32{9, 1}
	//	firstList = createSinglyLinkedList(firstArray)
	//	reverseList(&firstList)
	//	//printList(firstList)
	//
	//	//CREATE DOUBLY LINKED LIST
	//
	//	firstArray = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//	firstList = createSinglyLinkedList(firstArray)
	//	doublyListHead := createDoublyLLFromSinlgyLL(firstList)
	//
	//	doublyListHead.printList()
	//	//printList(secondList)
	//	//Append a new element
	//	// newElement := int32(19)
	//	// appendToEnd(head, newElement)
	//	// anotherElement := int32(32)
	//	// head = appendToBeginning(head, anotherElement)
	//
	//	// Delete the first element
	//	// head = deleteFirstElement(head)
	//
	//	//Put last element in head
	//	//head = moveLastToBeginning(head)
	//	//head = moveFistToLast(head)
	//	// reverse linkedlist
	//	//	head = reverseList(head)
	//	//	printList(head)

}

// reverse linked list in O(1) space
//		 1 > 2 > 3 > 4 > 5 > 6
//		 1 < 2
//		 6 > 5 > 4 > 3 > 2 > 1

func deleteList(head **singlyLinkedListNode) {
	*head = nil
}

func moveFistToLast(head *singlyLinkedListNode) *singlyLinkedListNode {

	lastElement := head
	for ; lastElement.next != nil; lastElement = lastElement.next {

	}

	lastElement.next = head

	tmp := head.next
	head.next = nil
	head = tmp

	return tmp

}

func moveLastToBeginning(head *singlyLinkedListNode) *singlyLinkedListNode {
	// Get to the previous of last element
	currentHead := head
	for ; currentHead.next.next != nil; currentHead = currentHead.next {
	}
	var tmp *singlyLinkedListNode
	tmp = currentHead.next
	currentHead.next = nil
	tmp.next = head
	return tmp
}

func appendToEnd(head *singlyLinkedListNode, newElement int32) {
	p := head
	for ; p.next != nil; p = p.next {
	}
	p.next = &singlyLinkedListNode{data: newElement}
}

func appendToBeginning(head *singlyLinkedListNode, newElement int32) *singlyLinkedListNode {
	newNode := singlyLinkedListNode{data: newElement, next: head}
	head = &newNode
	return head
}

func (list doublyLinkedListNode) printList() {
	for p := &list; p != nil; p = p.next {
		fmt.Println(p.data)
	}

}

func (list singlyLinkedListNode) printList() {
	for p := &list; p != nil; p = p.next {
		fmt.Println(p.data)
	}

}

func deleteFirstElement(head *singlyLinkedListNode) *singlyLinkedListNode {
	return head.next
}

func createSinglyLinkedList(arr []int32) *singlyLinkedListNode {
	initialNode := singlyLinkedListNode{
		data: arr[0],
		next: nil,
	}

	head := &initialNode
	builder := &initialNode

	for i := 1; i < len(arr); i++ {
		newNode := singlyLinkedListNode{
			data: arr[i],
			next: nil,
		}
		builder.next = &newNode
		builder = &newNode
	}
	return head
}

func createDoublyLinkedList(arr []int32) *doublyLinkedListNode {
	head := &doublyLinkedListNode{
		data: arr[0],
		next: nil,
	}

	builder := head

	for i := 1; i < len(arr); i++ {
		newNode := &doublyLinkedListNode{
			data: arr[i],
			next: nil,
			prev: builder,
		}
		builder.next = newNode
		builder = newNode
	}
	return head
}

func getNodeFromTail(head linkedList, positionFromTail int32) int32 {
	currentHead := head
	elementsInList := count(head)
	nthElement := elementsInList - positionFromTail
	for i := int32(0); i < nthElement; i++ {
		currentHead = *(currentHead.getNext())
	}
	return currentHead.getData().(int32)
}

func count(head linkedList) int32 {
	counter := int32(0)
	currentHead := head
	for currentHead != nil {
		counter++
		currentHead = *(currentHead.getNext())
	}
	return counter
}

func pop(head *singlyLinkedListNode) {
	head = head.next
}

func push(head *singlyLinkedListNode, element int32) {
	newNode := &singlyLinkedListNode{data: element, next: head.next}
	head = newNode
}

func getNode(head linkedList, index int32) int32 {
	currentHead := head
	for i := int32(0); i < index; i++ {
		currentHead = *(currentHead.getNext())
	}
	return currentHead.getData().(int32)
}

func insertNth(head *singlyLinkedListNode, index int, element int32) {
	currentHead := head
	for i := 0; i < index-1 && currentHead.next != nil; i++ {
		currentHead = currentHead.next
	}

	newNode := &singlyLinkedListNode{data: element, next: currentHead.next}
	currentHead.next = newNode
}

func sortedInsert(head *singlyLinkedListNode, insertedElement int32) {
	currentHead := *head
	currentMaximum := currentHead.getData()

	newNode := singlyLinkedListNode{
		data: insertedElement,
	}

	//if already smaller than first, append to beginning
	if insertedElement < currentMaximum {
		newNode.next = *head
		*head = &newNode
		return
	}

	for currentHead.next != nil && insertedElement > currentHead.next.data {
		currentMaximum = currentHead.next.data
		currentHead = currentHead.next
	}
	newNode.next = currentHead.next
	currentHead.next = &newNode
}

func appendLinkedLists(firstList *singlyLinkedListNode, secondList *singlyLinkedListNode) {

	// get to the last element of firstList and make .next the head of secondList
	currentHead := firstList
	for currentHead.next != nil {
		currentHead = currentHead.next
	}

	currentHead.next = secondList

}

//modifies the initial head and returns the secondHead from the middle
func frontBackSplit(head linkedList) *linkedList {
	//COUNT ELEMENTS IN LIST
	//TRAVERSE TO ELEMENT IN CEIL(MIDDLE)
	//SECONDHEAD = *CEIL(MIDDLE).next
	//NODE CEIL(MIDDLE).next = nil
	slowHead := &head
	fastHead := &head
	for fastHead.getNext() != nil && (*fastHead.getNext()).getNext() != nil {
		slowHead = slowHead.getNext()
		fastHead = fastHead.next.next
	}
	if fastHead.next != nil {
		fastHead = fastHead.next
		slowHead = slowHead.next
	}
	secondHead := slowHead.next
	slowHead.next = nil
	return secondHead
}

//Removes duplicates from a sorted list
func removeDuplicates(head *singlyLinkedListNode) {
	currentHead := head
	for currentHead.next != nil {
		if currentHead.data == currentHead.next.data {
			currentHead.next = currentHead.next.next
		} else {
			currentHead = currentHead.next
		}
	}
}

// receives a sorted array with duplicates in it and returns an array with unique elements
func buildSetFromArray(arr []int32) []int32 {
	newArray := []int32{}
	uniqueElements := make(map[int32]bool)
	for _, element := range arr {
		if exists := uniqueElements[element]; !exists {
			newArray = append(newArray, element)
		}
		uniqueElements[element] = true
	}
	return newArray
}

func compareArrayWithLinkedList(arr []int32, head linkedList) bool {
	for i := 0; head != nil && i < len(arr); head, i = head.getNext(), i+1 {
		if head.data != arr[i] {
			areEqual = false
			log.Println(head.data, " !=", arr[i])
		}
	}
	return true
}

func moveNode(firstHead *singlyLinkedListNode, secondHead *singlyLinkedListNode) {
	oldSecondHead := secondHead.next
	secondHead.next = firstHead
	firstHead = secondHead
	secondHead = oldSecondHead
}

func alternatingSplit(head *linkedList) *linkedList {

	generatedHead := head.next
	secondHead := generatedHead
	for currentHead := head; currentHead.next != nil; {
		currentHead.next = currentHead.next.next
		currentHead = currentHead.next

		generatedHead.next = generatedHead.next.next
		generatedHead = generatedHead.next
	}
	return secondHead
}

//MERGES BOTH LISTS IN THE FIRST LIST, DELETING THE SECOND ONE
func shuffleMerge(firstList linkedList, secondList *linkedList) {

	for (*secondList) != nil && firstList != nil {
		secondListFirst, secondListSecond := *secondList, (*secondList).next
		secondListFirst.next = firstList.next
		firstList.next = secondListFirst
		*secondList = secondListSecond
		firstList = firstList.next.next

	}

}

//INSERT THE SECOND LIST INTO THE FIRST LIST SO THAT THE FIRST LIST IS ORDERED
func sortedMerge(firstHead *linkedList, secondHead *linkedList) *linkedList {
	var newHead *singlyLinkedListNode

	//choose which element to start with, firstHead or secondHead, choose the smallestone
	if (*firstHead).data < (*secondHead).data {
		newHead = *firstHead
		(*firstHead) = (*firstHead).next
	} else {
		newHead = *secondHead
		(*secondHead) = (*secondHead).next
	}

	movingHead := newHead

	for (*firstHead) != nil && (*secondHead) != nil {
		if (*firstHead).data < (*secondHead).data {
			movingHead.next = (*firstHead)
			*firstHead = (*firstHead).next
		} else {
			movingHead.next = *secondHead
			*secondHead = (*secondHead).next
		}
		movingHead = movingHead.next
	}

	if *firstHead == nil {
		movingHead.next = *secondHead
	} else if *secondHead == nil {
		movingHead.next = *firstHead
	}
	return newHead
}

func mergeSort(head *linkedList) {
	//split the list into smaller lists
	if (*head) == nil || count(*head) == 1 {
		return
	}
	lengthOfList := count(*head)
	newHead := *head
	for i := int32(0); i < int32(lengthOfList/2)-1; i++ {
		newHead = newHead.next
	}
	tmp := newHead.next
	newHead.next = nil
	newHead = tmp

	mergeSort(head)
	mergeSort(&newHead)

	var smallestHead *singlyLinkedListNode
	if newHead.data < (*head).data {
		smallestHead = newHead
		newHead = newHead.next
	} else {
		smallestHead = *head
		*head = (*head).next
	}
	initialNewHead := smallestHead
	for newHead != nil && *head != nil {
		if newHead.data < (*head).data {
			smallestHead.next = newHead
			newHead = newHead.next
		} else {
			smallestHead.next = *head
			*head = (*head).next
		}
		smallestHead = smallestHead.next
	}

	if *head == nil {
		smallestHead.next = newHead
	} else {
		smallestHead.next = *head
	}
	*head = initialNewHead
}

//findCommonInArray also works with unsorted arrays
func findCommonInArray(firstArray []int32, secondArray []int32) []int32 {

	var bigArray *[]int32
	var smallArray *[]int32
	if len(firstArray) > len(secondArray) {
		bigArray = &firstArray
		smallArray = &secondArray
	} else {
		bigArray = &secondArray
		smallArray = &firstArray
	}

	initialdDictionary := make(map[int32]bool)
	commonElements := []int32{}
	for _, v := range *bigArray {
		initialdDictionary[v] = true
	}

	for _, v := range *smallArray {
		if initialdDictionary[v] {
			commonElements = append(commonElements, v)
		}
	}
	return commonElements
}

func sortedIntersect(head1 *singlyLinkedListNode, head2 *singlyLinkedListNode) *singlyLinkedListNode {
	movingHead1 := head1
	movingHead2 := head2

	var pointerToReturn *singlyLinkedListNode
	var initialNode *singlyLinkedListNode
	for movingHead1 != nil && movingHead2 != nil {
		if movingHead1.data == movingHead2.data {
			if pointerToReturn == nil {
				pointerToReturn = &singlyLinkedListNode{data: movingHead1.data}
				initialNode = pointerToReturn
			} else {
				pointerToReturn.next = &singlyLinkedListNode{data: movingHead1.data}
				pointerToReturn = pointerToReturn.next
			}
			movingHead1 = movingHead1.next
			movingHead2 = movingHead2.next
		} else if movingHead1.data < movingHead2.data {
			movingHead1 = movingHead1.next
		} else {
			movingHead2 = movingHead2.next
		}
	}

	return initialNode
}

func reverseList(head *singlyLinkedListNode) {
	newHead := &singlyLinkedListNode{}

	for head != nil {
		newNode := head
		newNode.next = newHead.next
		newHead.next = newNode
		head = head.next
	}
	head = newHead
}
