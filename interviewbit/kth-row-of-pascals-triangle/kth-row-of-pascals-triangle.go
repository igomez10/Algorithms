package main

import "fmt"

func main() {
	//	initialRow := []int{1, 5, 10, 10, 5, 1}
	//	expectedRow := []int{1, 6, 15, 20, 15, 6, 1}
	//	nextRow := GenerateNextRow(initialRow)
	//	fmt.Println(nextRow)
	//	for i := 0; i < len(nextRow); i++ {
	//		if nextRow[i] != expectedRow[i] {
	//			fmt.Println("FAILED")
	//		} else {
	//			fmt.Println("SUCCESS")
	//		}
	//	}
	fmt.Println(getRow(0))
}

func getRow(level int) []int {
	if level == 0 {
		return []int{1}
	} else if level == 1 {
		return GenerateNextRow([]int{1})
	}
	return GenerateNextRow(getRow(level - 1))
}

// GenerateNextRow returns the next pascal pascal triangle given row
// assumes the initial row is valid
func GenerateNextRow(initialRow []int) []int {
	newArray := make([]int, len(initialRow)+1)

	newArray[0] = 1
	newArray[len(newArray)-1] = 1
	for i := 1; i < len(newArray)-1; i++ {
		newArray[i] = initialRow[i] + initialRow[i-1]
	}
	return newArray
}
