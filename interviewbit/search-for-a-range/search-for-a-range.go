package main

import "fmt"

/*
Given a sorted array of integers, find the starting and ending position of a given target value.

Your algorithm’s runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].


*/

func main() {
	arr := []int{5, 7, 7, 8, 8, 10}
	target := 8
	foundIndex := binarySearch(arr, target, 0, len(arr)-1)
	fmt.Println("foundIndex", foundIndex)
	fmt.Println(findFirst(arr, foundIndex))
	fmt.Println(findLast(arr, foundIndex))
}

func solve() {

}

func binarySearch(arr []int, target, start, end int) int {
	var middle int

	middle = (end + start) / 2

	for middle >= 0 && middle < len(arr) && start <= end {
		if arr[middle] == target {
			return middle
		} else if arr[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
		middle = (end + start) / 2
	}
	return -1
}

func findFirst(arr []int, foundIndex int) int {
	originalValue := arr[foundIndex]
	firstElement := foundIndex
	for ; firstElement >= 0; foundIndex-- {
		if arr[foundIndex] != originalValue {
			firstElement = foundIndex
			break
		}
	}
	return firstElement + 1
}

func findLast(arr []int, foundIndex int) int {
	originalValue := arr[foundIndex]
	lastElement := foundIndex
	for ; lastElement < len(arr); foundIndex++ {
		if arr[foundIndex] != originalValue {
			lastElement = foundIndex
			break
		}

	}
	return lastElement - 1
}
