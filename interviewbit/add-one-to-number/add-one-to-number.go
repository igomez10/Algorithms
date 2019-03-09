package main

import "fmt"

func main() {
	arr := []int{0, 0, 0, 0}
	arr = addOneToNumber(arr)
	fmt.Println(arr)
}

func addOneToNumber(arr []int) []int {
	var remainder int
	newValue := (arr)[len(arr)-1] + 1

	if newValue == 10 {
		remainder++
		newValue = 0
	}

	(arr)[len(arr)-1] = newValue
	i := len(arr) - 2
	for ; i >= 0; i-- {
		(arr)[i] += remainder
		remainder = 0
		if (arr)[i] == 10 {
			(arr)[i] = 0
			remainder++
		} else {
			break
		}

	}
	firstNum := 0
	isNull := true
	for ; firstNum < len(arr); firstNum++ {
		if arr[firstNum] != 0 {
			isNull = false
			break
		}
	}
	if !isNull {
		newArray := make([]int, len(arr)-firstNum+1)
		newArray = arr[firstNum:]
		arr = newArray
	} else {

		arr = make([]int, len(arr)+1)
		arr[0] = 1
	}

	return arr
}
