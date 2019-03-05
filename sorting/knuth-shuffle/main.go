package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(shuffle(arr))
}

func shuffle(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		source := rand.NewSource(time.Now().UnixNano())
		generator := rand.New(source)
		rIndex := generator.Intn(len(arr))
		arr[rIndex], arr[i] = arr[i], arr[rIndex]
	}
	return arr
}
