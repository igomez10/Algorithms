package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type arrLargest []int

// Less rerturn true if the element in position A is a smaller addition
// than the element in position B
func (arr arrLargest) Less(A, B int) bool {
	optionA := strconv.Itoa(arr[A]) + strconv.Itoa(arr[B])
	optionB := strconv.Itoa(arr[B]) + strconv.Itoa(arr[A])
	parsedA, _ := strconv.Atoi(optionA)
	parsedB, _ := strconv.Atoi(optionB)

	if parsedA < parsedB {
		return true
	}
	return false
}

func (arr arrLargest) Swap(A, B int) {
	arr[A], arr[B] = arr[B], arr[A]
}

func (arr arrLargest) Len() int {
	return len(arr)
}

func main() {
	fmt.Println("vim-go")
}

func findLargest(arr []int) string {
	sortableArr := arrLargest(arr)
	builder := strings.Builder{}
	sort.Sort(sortableArr)

	for i := len(arr) - 1; i >= 0; i-- {
		builder.WriteString(strconv.Itoa(sortableArr[i]))
	}

	return builder.String()

}
