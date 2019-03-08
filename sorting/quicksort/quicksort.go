package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 8, 7, 6, 4, 3, 2, 1}

	sort(&arr, 0, len(arr)-1)

	fmt.Println(arr)
	fmt.Println("EOF")
}

func sort(arr *[]int, lo, hi int) {
	if lo >= hi {
		return
	}
	sortedIndex := partition(arr, lo, hi)
	partition(arr, lo, sortedIndex)
	partition(arr, sortedIndex, hi)
}

func partition(arr *[]int, lo, hi int) int {
	if len(*arr) <= 1 || lo >= hi {
		return 0
	}
	pivotElement := (*arr)[0]
	for {
		for (*arr)[lo] < pivotElement && lo <= hi {
			lo++
		}
		for (*arr)[hi] > pivotElement && hi >= lo {
			hi--
		}
		if hi <= lo {
			break
		} else {
			(*arr)[lo], (*arr)[hi] = (*arr)[hi], (*arr)[lo]
		}
	}
	(*arr)[hi], (*arr)[0] = (*arr)[0], (*arr)[hi]
	return hi
}
