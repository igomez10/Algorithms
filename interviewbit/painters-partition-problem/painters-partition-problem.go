package main

import "fmt"

/*
You have to paint N boards of length {A0, A1, A2, A3 … AN-1}.

There are K painters available and you are also given how much time a
painter takes to paint 1 unit of board.

You have to get this job done as soon as possible under the constraints that any
painter will only paint contiguous sections of board.
*/

/**
 * @input A : Integer
 * @input B : Integer
 * @input C : Integer array
 *
 * @Output Integer
 */
func mmax(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func possible(numPainters, timeToPaint int, arr []int, mid int64) bool {
	s := int64(0)
	for i := 0; i < len(arr); i++ {
		if s+int64(arr[i]*timeToPaint) <= mid {
			s += int64(arr[i] * timeToPaint)
		} else {
			numPainters--
			if numPainters == 0 {
				return false
			}
			s = int64(arr[i] * timeToPaint)
			if s > mid {
				return false
			}
		}
	}
	return s <= mid
}

func paint(numPainters int, timeToPaint int, boards []int) int {
	max := int64(0)
	last := int64(0)
	// get longest board to pain with a single iteration over painters
	for i := 0; i < numPainters-1; i++ {
		max = mmax(int64(boards[i]*timeToPaint), max)
		if i == len(boards)-1 {
			break
		}
	}
	// how to paint the remaining boards with the previous painters
	// add it to max
	for i := numPainters - 1; i < len(boards); i++ {
		max += int64(boards[i] * timeToPaint)
	}
	min := int64(0)
	var mid int64
	for min <= max {
		mid = (max + min) / 2
		if possible(numPainters, timeToPaint, boards, mid) {
			last = mid
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return int(last % 10000003)
}

func main() {
	K := 2
	T := 5
	L := []int{1, 10}
	fmt.Println(paint(K, T, L))
}
