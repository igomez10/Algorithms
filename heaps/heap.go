package main

import "fmt"
import "math"
import "strings"

/*
Build max heap from unordered array
Find max element a[i]
swap elements a[n] a[i] now max element is at the end of array
discard node n from heap, decrementing heap size
new root after the swap may violate max heap property but its children are max heaps, we can run max heapify to fix this
*/

func main() {
	values := []int{-1, 9, 7, 6, 5, 4, 3, 2, 1}
	printHeap(values)
	heapSort(&values)
	fmt.Println()
	printHeap(values)
	//	fmt.Println("Initial Array:")
	//	printHeap(values)
	//	inserted := 2
	//	insert(&values, inserted)
	//	fmt.Printf("After inserting %d:\n", inserted)
	//	printHeap(values)
	//	fmt.Printf("Heap is sorted: %t\n", isSorted(values))
}

func heapSort(arr *[]int) {
	var i int
	for i = len(*arr) / 2; i >= 1; i-- {
		fmt.Println("sinking", (*arr)[i])
		sink(arr, i)
	}
	fmt.Println(arr)
	//	for i := 1; i < len(*arr); i++ {
	//		insert(&newHeap, (*arr)[i])
	//	}
	//	sortedHead := []int{-1}
	//	for i := 1; i < len(*arr); i++ {
	//		sortedHead = append(sortedHead, newHeap[1])
	//		deleteMin(&newHeap)
	//	}

}

func deleteMin(arr *[]int) {
	// overwrite last element to first position in heap
	(*arr)[1] = (*arr)[len(*arr)-1]
	// remove last element
	*arr = (*arr)[:len(*arr)-1]
	// move first element to its final position
	sink(arr, 1)
}

func printHeap(heap []int) {
	arr := heap[1:]
	levelsInArr := int(math.Ceil(math.Log2(float64(len(arr)))))
	if len(heap) <= 3 {
		levelsInArr = 1
	}
	for i := 0; i <= levelsInArr; i++ {
		var row []int
		index := float64(i)
		numElementsInRow := int(math.Exp2(index))

		if numElementsInRow <= len(arr) {
			row, arr = arr[:numElementsInRow], arr[numElementsInRow:]
		} else {
			row = arr
		}

		initalSpace := strings.Repeat(" ", levelsInArr-i)
		if len(row) > 0 {
			fmt.Println(initalSpace, row)
		} else {
			break
		}
	}
}

func isSorted(arr []int) bool {
	isSorted := true
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i/2] {
			isSorted = false
			break
		}
	}

	return isSorted
}

func isALeaf(arr *[]int, position int) bool {
	if len(*arr) < position*2 {
		return true
	}
	return false
}

func hasRightChild(arr *[]int, position int) bool {
	if len(*arr) < position*2+1 {
		return true
	}
	return false
}

func isSmallerThanChilds(arr *[]int, position int) bool {
	if isALeaf(arr, position) {
		return true
	}

	var answer bool
	//compare with left child, then with right child
	if (*arr)[position*2] >= (*arr)[position] {
		// cannot be swapped with left child
		// compare with right child
		if hasRightChild(arr, position) {
			if (*arr)[position*2+1] >= (*arr)[position] {
				// is smaller than left AND right
				answer = true
			}
			// is smaller than left but bigger than right
			answer = false
		} else {
			// is smaller than left, doesnt have right
			answer = true
		}
	} else {
		// bigger than left
		answer = false
	}

	return answer
}

func sink(arr *[]int, position int) {
	current := position
	for {
		//replace left and right values in function
		left := current * 2
		right := current*2 + 1
		if isALeaf(arr, current) {
			break
		}
		var canSwapLeft bool
		var canSwapRight bool

		if (*arr)[current*2] < (*arr)[current] {
			// left child is smaller
			canSwapLeft = true
		}

		// check is element has right child
		if hasRightChild(arr, current) {
			if (*arr)[right] < (*arr)[current] {
				// right child is smaller
				canSwapRight = true
			}
		}
		if canSwapLeft && canSwapRight {
			// left and right child are smaller than current position
			// select the smallest to swap with current
			var smaller int
			if (*arr)[left] < (*arr)[right] {
				smaller = (*arr)[left]
			} else {
				smaller = (*arr)[right]
			}
			(*arr)[current], (*arr)[smaller] = (*arr)[smaller], (*arr)[current]
		} else if canSwapLeft {
			(*arr)[current], (*arr)[left] = (*arr)[left], (*arr)[current]
			current = left
		} else if canSwapRight {
			(*arr)[current], (*arr)[right] = (*arr)[right], (*arr)[current]
			current = right
		} else {
			break
		}
	}
}

// sink takes the given index and moves
//func sink(arr *[]int, indexOfValue int) int {
//	// if a value is in last row do nothing
//	i := indexOfValue
//	// the element is in its final position if its a leaf or if its smaller than both childs
//	isInFinalPosition := isALeaf(arr, i) || (((*arr)[i] <= (*arr)[i*2]) && (*arr)[i] <= (*arr)[i*2+1])
//	for !isInFinalPosition {
//		left := (*arr)[i*2]
//		right := (*arr)[i*2+1]
//		var smallest int
//		if right > left {
//			smallest = i * 2
//		} else {
//			smallest = i*2 + 1
//		}
//		(*arr)[i], (*arr)[smallest] = (*arr)[smallest], (*arr)[i]
//		i = smallest
//
//		isInFinalPosition = isALeaf(arr, i) || (((*arr)[i] <= (*arr)[i*2]) && (*arr)[i] <= (*arr)[i*2+1])
//	}
//
//	return i
//}

func isInHeap(arr []int, elementToFind int) bool {
	i := 0
	for i < len(arr) {
		if arr[i] == elementToFind {
			return true
		} else if arr[i*2] < elementToFind {
			i = i*2 + 1
		} else {
			i = i * 2
		}
	}
	return false
}

func swim(arr *[]int, indexOfValue int) {
	i := indexOfValue
	//if index is root or if parent is smaller than Value-> do nothing

	isRoot := i == 0
	valueIsInFinalPosition := isRoot || (*arr)[i/2] <= (*arr)[i]
	for !valueIsInFinalPosition {
		(*arr)[i], (*arr)[i/2] = (*arr)[i/2], (*arr)[i]
		i = i / 2
		isRoot = i == 0
		valueIsInFinalPosition = isRoot || (*arr)[i/2] < (*arr)[i]
	}
}

func insert(arr *[]int, newValue int) {
	*arr = append(*arr, newValue)
	swim(arr, len(*arr)-1)
}
