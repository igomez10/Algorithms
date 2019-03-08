package main

import "fmt"
import "time"

func main() {
	arr := []int{1, 2, 3, 10, 45, 67, 89, 100, 125, 167, 200, 245, 255, 267, 290, 340, 350, 500, 900, 9000, 100000, 10000000}

	fmt.Println(binarySearch(&arr, 0, len(arr)-1, 255))
}

func binarySearch(arr *[]int, lo, hi, searched int) bool {

	time.Sleep(1 * time.Second)

	fmt.Printf("From %d (%d) to %d (%d) \n", lo, (*arr)[lo], hi, (*arr)[hi])
	if hi == lo || lo > hi {
		return false
	}

	var answer bool
	middleIndex := (hi + lo) / 2
	pivotElement := (*arr)[middleIndex]

	if pivotElement == searched {
		answer = true
	} else if pivotElement > searched {
		answer = binarySearch(arr, lo, middleIndex, searched)
	} else {
		answer = binarySearch(arr, middleIndex, hi, searched)
	}
	return answer

}
