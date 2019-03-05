package main

func newnewSelectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		minimumIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minimumIndex] {
				minimumIndex = j
			}
		}
		arr[i], arr[minimumIndex] = arr[minimumIndex], arr[i]
	}

	return arr
}
