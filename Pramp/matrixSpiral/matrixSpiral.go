package main

import "fmt"

func SpiralCopy(inputMatrix [][]int) []int {
	up := 0
	down := len(inputMatrix) - 1

	left := 0
	right := len(inputMatrix[0]) - 1

	ans := []int{}
	// i := 0
	// j := 0
	// direction := "right"
	for len(ans) < len(inputMatrix)*len(inputMatrix[0]) {

		for i := left; i <= right; i++ {
			ans = append(ans, inputMatrix[up][i])
		}
		up++

		for i := up; i <= down; i++ {
			ans = append(ans, inputMatrix[i][right])
		}
		right--

		for i := right; i >= left; i-- {
			ans = append(ans, inputMatrix[down][i])
		}
		down--

		for i := down; i >= up; i-- {
			ans = append(ans, inputMatrix[i][left])
		}
		left++
	}

	return ans
}

/*
[1,    2,   3,  4,    5],
[6,    7,   8,  9,   10],
[11,  12,  13,  14,  15],
[16,  17,  18,  19,  20] ]

[1,    0,   0,  0,    0],
[0,    0,   0,  0,    0],
[0,    0,   0,  0,    0],
[0,    0,   0,  0,    0],

up = 0
down = len(arr)

left = 0
right = len(arr[0])

arr[up][left] -> arr[up][right]
arr[up][right] -> arr[down][right]
right--
up++
arr[down][right] -> arr[down][left]
arr[down][left] -> arr[up][left]
down--
left++

go to right
go down
go left
go up






*/

func main() {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	fmt.Println(SpiralCopy(input))
}
