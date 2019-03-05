package main

import "fmt"
import "sort"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//hey := "Hello world my name is Ignacio"
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
}
