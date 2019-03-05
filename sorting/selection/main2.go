package main

func newSelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		indexMinimum := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[indexMinimum] {
				indexMinimum = j
			}
			arr[i], arr[indexMinimum] = arr[indexMinimum], arr[i]
		}
	}
	return arr
}
