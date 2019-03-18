package main

import "fmt"

func main() {
	arr := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	rotateMatrix(&arr)
	fmt.Println(arr)
}

func rotateMatrix(arr *[][]int) {
	last := len(*arr) - 1
	for i := 1; i < len(*arr); i++ {
		tmp := (*arr)[0][i]
		(*arr)[i][last], tmp = tmp, (*arr)[i][last]
		(*arr)[last][last-i], tmp = tmp, (*arr)[last][last-i]
		(*arr)[last-i][0], tmp = tmp, (*arr)[last-i][0]
		(*arr)[0][i] = tmp
	}
}
