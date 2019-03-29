package main

import "fmt"
import "math"

func main() {
	fmt.Println("vim-go")
}

func isOneAway(A, B string) bool {
	if math.Abs(float64(len(A)-len(B))) > 1 {
		return false // if the difference in len is bigger than 1 return false
	}
	var small, big string
	if len(A) < len(B) { // select the big and small string
		small = A
		big = B
	} else {
		small = B
		big = A
	}
	var index1, index2 int
	var foundDifference bool
	for index1 < len(small) && index2 < len(big) {

		if small[index1] != big[index2] { // difference in current index

			if foundDifference { //if a difference was found previously
				return false
			}
			foundDifference = true // save this found difference for future

			// if we got here this is the first found difference, increse index in small by 1
			if len(small) == len(big) {
				index1++ //only move the small index if this was caused by a replacement
			}
		} else {
			index1++
		}
		index2++
	}
	return true
}

func customIsOneAway(A, B string) bool {
	if math.Abs(float64(len(A)-len(B))) > 1 {
		return false
	}

	var bigger string
	var smaller string
	if len(A) > len(B) {
		bigger = A
		smaller = B
	} else {
		bigger = B
		smaller = A
	}

	var foundDifference bool
	bigPointer := 0
	smallPointer := 0
	for smallPointer < len(smaller) && bigPointer < len(bigger) {
		if bigger[bigPointer] != smaller[smallPointer] {
			if foundDifference {
				return false
			}
			foundDifference = true
			if len(bigger) == len(smaller) {
				smallPointer++
			}
		} else {
			smallPointer++
		}

		bigPointer++
	}
	return true
}
