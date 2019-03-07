package main

import "fmt"

func main() {
	arr := [][]int32{
		[]int32{-9, -9, -9, 1, 1, 1},
		[]int32{0, -9, 0, 4, 3, 2},
		[]int32{-9, -9, -9, 1, 2, 3},
		[]int32{0, 0, 8, 6, 6, 0},
		[]int32{0, 0, 0, -2, 0, 0},
		[]int32{0, 0, 1, 2, 4, 0},
	}
	maxCount := int32(-9 * 7)
	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr)-2; j++ {
			currentSum := int32(0)
			//upper row
			currentSum += arr[i][j]
			currentSum += arr[i][j+1]
			currentSum += arr[i][j+2]

			// middle num
			currentSum += arr[i+1][j+1]

			//lower row
			currentSum += arr[i+2][j]
			currentSum += arr[i+2][j+1]
			currentSum += arr[i+2][j+2]
			if currentSum > maxCount {
				maxCount = currentSum
			}
		}

	}
	fmt.Println(maxCount)
}
