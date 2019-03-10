package main

import "fmt"
import "math"
import "strings"

func main() {
	values := []int{1, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Initial Array:")
	printHeap(values)
	inserted := 2
	insert(&values, inserted)
	fmt.Printf("After inserting %d:\n", inserted)
	printHeap(values)
	fmt.Printf("Heap is sorted: %t", isSorted(values))
}

func printHeap(arr []int) {

	levelsInArr := int(math.Log2(float64(len(arr))))
	for i := 0; i <= levelsInArr; i++ {
		var row []int
		index := float64(i)
		numElementsInRow := int(math.Exp2(index))

		if numElementsInRow <= len(arr) {
			row, arr = arr[:numElementsInRow], arr[numElementsInRow:]
		} else {
			row = arr
		}

		initalSpace := strings.Repeat(" ", (levelsInArr*levelsInArr)/2-i)
		fmt.Println(initalSpace, row)
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

func swim(arr *[]int, indexOfValue int) {
	i := indexOfValue

	//if index is root or if parent is smaller than Value-> do nothing
	valueIsInFinalPosition := i == 0 || (*arr)[i/2] < (*arr)[i]
	if !valueIsInFinalPosition {
		(*arr)[i], (*arr)[i/2] = (*arr)[i/2], (*arr)[i]
		swim(arr, i/2)
	}
}

func insert(arr *[]int, newValue int) {
	*arr = append(*arr, newValue)
	//parent := (*arr)[len(*arr)/2]
	if ok := isSorted(*arr); !ok {
		swim(arr, len(*arr)-1)
	}
}
