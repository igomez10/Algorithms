package main

import "fmt"
import "strconv"
import "strings"

func main() {
	fmt.Println("vim-go")
}

func RotateMatrix(arr *[][]int) {

	for i := 0; i < len(*arr)-1; i++ {
		tmp := (*arr)[0][i]

		(*arr)[i][len(*arr)-1], tmp = tmp, (*arr)[i][len(*arr)-1]

		(*arr)[len(*arr)-1][len(*arr)-1-i], tmp = tmp, (*arr)[len(*arr)-1][len(*arr)-1-i]

		(*arr)[len(*arr)-1-i][0], tmp = tmp, (*arr)[len(*arr)-1-i][0]

		(*arr)[0][i], tmp = tmp, (*arr)[0][i]
	}
}
func PrintMatrix(arr [][]int) {
	for i := range arr {
		for j := range arr {
			fmt.Println(arr[i][j])
		}
	}
}
func matrixToString(arr [][]int) string {
	builder := strings.Builder{}
	for i := range arr {
		for j := range arr {
			builder.WriteString(strconv.Itoa(arr[i][j]))
			builder.WriteString(" ")
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
