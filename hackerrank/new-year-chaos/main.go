package main

import "fmt"
import "math"

// 1,2,3,4,5
// 2,1,3,4,5
// 2,3,1,4,5
// 2,3,4,1,5

//1,2,3,4,5,6,7,8
//1,2,3,5,4,6,7,8
//1,2,5,3,4,6,7,8
//1,2,5,3,4,7,6,8
//1,2,5,3,7,4,6,8
//1,2,5,3,7,4,8,6
//1,2,5,3,7,8,4,6
//1,2,5,3,7,8,6,4
//| | | | | | | |
//1,2,3,4,5,6,7,8
//| | | | | | | |
//0,0,2,0,2,2,0,0
func main() {

	arr1 := []int32{2, 1, 5, 3, 4}
	arr2 := []int32{2, 5, 1, 3, 4}

	arr3 := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	arr4 := []int32{2, 3, 4, 5, 6, 1}
	testCases := [][]int32{arr1, arr2, arr3, arr4}
	for v := range testCases {
		bribesinI, isPossibleI := countBribesSolution(testCases[v])
		if !isPossibleI {
			fmt.Printf("Too chaotic: %+v\n", testCases[v])
		} else {
			fmt.Println(bribesinI)
		}

	}

}

func countBribes(arr []int32) (int, bool) {
	isPossible := true
	counter := 0
	for i := 0; i < len(arr); i++ {
		bribesforI := (int(arr[i]) - 1) - i
		if bribesforI > 2 {
			isPossible = false
			break
		} else if bribesforI > 0 {
			counter += int(bribesforI)
		}
	}

	return counter, isPossible
}

func countBribesSolution(arr []int32) (int, bool) {
	isPossible := true
	counter := 0
	for i := len(arr) - 1; i >= 0; i-- {
		if int(arr[i])-(i+1) > 2 {
			isPossible = false
			break
		}
		j := int(math.Max(0, float64(arr[i]-2)))
		for ; j < 1; j++ {
			if arr[j] > arr[i] {
				counter++
			}
		}
	}
	return counter, isPossible
}
