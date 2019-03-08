package main

import "fmt"

func main() {
	arr := []int{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 200, 300, 400, 555, 600, 700, 800, 900, 1000}
	searchedElement := 1005

	//testing all the values in the array
	for i := range arr {
		if !isInArray(arr, arr[i]) {
			fmt.Println("Not found:", arr[i])
		}
	}

	fmt.Println(isInArray(arr, searchedElement))
}

func isInArray(arr []int, searchedElement int) bool {

	lo := 0
	hi := len(arr) - 1

	//variable to be returned in the end, zero value is false
	var answer bool

	//it's ok if the indexes are equal, after the next cycle, lo|hi will move +-1, preventing constant values as lo==hi
	for lo <= hi {
		middle := (hi + lo) / 2
		pivotElement := arr[middle]
		if pivotElement == searchedElement {
			answer = true
			break
		} else if pivotElement < searchedElement {
			lo = middle + 1 // why include the middle element if we already know that it's not our searchedElement??, skip it with a +1
		} else { // same here
			hi = middle - 1 // same here
		}
	}

	return answer

}
