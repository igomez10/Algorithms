package main

type llnode struct {
	value int
	next  *llnode
}

func main() {}

func count(listA llnode) int {
	var counter int
	for listA.next != nil {
		counter++
		listA = *listA.next
	}
	return counter
}

func sumLists(listA, listB *llnode) *llnode {

	longList := listA
	shortList := listB

	if count(*longList) < count(*shortList) {
		shortList = listA
		longList = listB
	}

	originalHead := longList
	carryOn := 0

	for shortList != nil {
		longList.value += shortList.value

		if carryOn > 0 {
			longList.value += carryOn
			carryOn = 0
		}

		if longList.value >= 10 {
			carryOn = longList.value / 10
			longList.value = longList.value % 10
		}

		shortList = shortList.next
		longList = longList.next
	}

	return originalHead
}
