package main

import "fmt"

func main() {
	arr := []int{73, 90, 85, 58, 69, 77, 90, 85, 18, 35}

	//	arr := []int{17, 74, 90, 85, 58, 69, 77, 90, 84, 36}

	fmt.Println(partition(arr, 0, len(arr)-1))
	fmt.Println(quickSelect(arr, 1))
}

func quickSelect(arr []int, k int) int {
	lo := 0
	hi := len(arr) - 1
	for {
		fmt.Println(arr)
		num := partition(arr, lo, hi)
		if num < k-1 {
			lo = num + 1
		} else if num > k-1 {
			hi = num - 1
		} else {
			return arr[num]
		}
	}
}

func partition(arr []int, y, z int) int {

	lo := y + 1
	hi := z
	pivot := arr[y]
	for lo < hi {
		for arr[lo] < pivot && lo < z {
			lo++
		}
		for arr[hi] > pivot && hi > y {
			hi--
		}
		if lo >= hi {
			break
		}
		arr[lo], arr[hi] = arr[hi], arr[lo]
	}
	arr[y], arr[hi] = arr[hi], arr[y]
	return hi
}
