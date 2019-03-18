package main

import "fmt"

/*
By using XOR we can implement a solution that can traverse the array
once and find which element is only once.
The XOR works by adding the number if its the first time its seen
and removing the array if it was already seen.
*/

func singleNumber(numbers []int) int {
	number := 0
	for i := 0; i < len(numbers); i++ {
		number ^= numbers[i]
	}

	return number
}

func main() {
	arr := []int{1, 2, 2, 3, 1}
	fmt.Println(singleNumber(arr))
}
