package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 2, 3, 4, 5, 6, 7, 5, 4, 3, 2, 2, 4, 6, 7, 8, 9}

	fmt.Println(newnewSelectionSort(arr))
}

func selectionSort(arr []int) []int {
	for j := 0; j < len(arr); j++ {
		for i := j; i < len(arr); i++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
