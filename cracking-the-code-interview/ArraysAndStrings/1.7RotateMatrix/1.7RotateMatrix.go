package main

import "fmt"
import "strconv"
import "strings"

func main() {
	fmt.Println("vim-go")
}

func RotateMatrix(arr *[][]int) {
	for initial := 0; initial < len(*arr)/2; initial++ {

		for i := initial; i < len(*arr)-initial-1; i++ {

			tmp := (*arr)[initial][i]

			(*arr)[i][len(*arr)-1-initial], tmp = tmp, (*arr)[i][len(*arr)-1-initial]

			(*arr)[len(*arr)-1-initial][len(*arr)-1-i], tmp = tmp, (*arr)[len(*arr)-1-initial][len(*arr)-1-i]

			(*arr)[len(*arr)-1-i][initial], tmp = tmp, (*arr)[len(*arr)-1-i][initial]

			(*arr)[initial][i], tmp = tmp, (*arr)[initial][i]
		}

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
