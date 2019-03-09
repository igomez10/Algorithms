package main

import "fmt"
import "math"

func main() {
	arr := []int{2, 2, 2}
	answer := maxArr(arr)
	fmt.Println(answer)
}

func maxArr(arr []int) int {
	ln := len(arr)

	mx1, mn1 := math.MinInt32, math.MaxInt32
	mx2, mn2 := math.MinInt32, math.MaxInt32

	for i := 0; i < ln; i++ {

		mx1 = max(mx1, arr[i]+i)
		mn1 = min(mn1, arr[i]+i)

		mx2 = max(mx2, arr[i]-i)
		mn2 = min(mn2, arr[i]-i)
	}

	mx := max(mx1-mn1, mx2-mn2)
	return mx
}

func min(A, B int) int {
	if B < A {
		return B
	}
	return A
}
func max(A, B int) int {
	if B > A {
		return B
	}
	return A
}
