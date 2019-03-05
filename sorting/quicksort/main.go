package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 2, 3, 4, 5, 6, 7, 5, 4, 3, 2, 2, 4, 6, 7, 8, 9}

	fmt.Println(quickSort(arr))

}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	smaller := []int{}
	bigger := []int{}
	pivot := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			smaller = append(smaller, arr[i])
		} else {
			bigger = append(bigger, arr[i])
		}
	}

	smaller = quickSort(smaller)
	bigger = quickSort(bigger)

	smaller = append(smaller, pivot)
	smaller = append(smaller, bigger...)
	return smaller
}
