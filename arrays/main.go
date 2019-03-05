package main

import "fmt"

func main() {
	arr := generateArray(5)
	arr = rotateImage(arr)

	for _, row := range arr {
		for _, element := range row {
			fmt.Print("     ", element)
		}
		fmt.Println()
	}

	fmt.Println()
}

func rotateImage(arr [][]int) [][]int {
	var numLevels int
	if len(arr) > 3 {
		numLevels = len(arr) - 2
	} else {
		numLevels = 1
	}
	for i := 0; i < numLevels; i++ {
		for j := i; j < len(arr)-1-i; j++ {
			tmp := arr[i][j]
			tmp, arr[j][len(arr)-i-1] = arr[j][len(arr)-i-1], tmp
			tmp, arr[len(arr)-i-1][len(arr)-j-1] = arr[len(arr)-i-1][len(arr)-j-1], tmp
			tmp, arr[len(arr)-j-1][i] = arr[len(arr)-j-1][i], tmp
			arr[i][j] = tmp
		}
	}
	return arr
}

func generateArray(numRows int) [][]int {

	newArray := make([][]int, numRows)
	counter := 1
	for i := range newArray {
		newArray[i] = make([]int, numRows)
		for j := range newArray {
			newArray[i][j] = counter
			counter++
		}
	}

	return newArray
}
