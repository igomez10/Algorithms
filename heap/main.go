package main

import "fmt"

func main() {

	h := heap{arr: []int{0}}

	itemsToInsert := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := range itemsToInsert {
		h.push(itemsToInsert[i])
	}

	for len(h.arr) > 1 {
		fmt.Println(h.pop())
	}
}

type heap struct {
	arr []int
}

func (h *heap) pop() int {
	ans := h.arr[1]
	h.arr[1] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	i := 1
	for i*2+1 <= len(h.arr)-1 && (h.arr[i] > h.arr[i*2] || h.arr[i] > h.arr[i*2+1]) {

		bestfitted := i * 2
		if h.arr[bestfitted] > h.arr[i*2+1] {
			bestfitted = i*2 + 1
		}

		h.arr[i], h.arr[bestfitted] = h.arr[bestfitted], h.arr[i]
		i = bestfitted

	}

	return ans

}

func (h *heap) push(n int) {

	h.arr = append(h.arr, n)
	finalIndex := len(h.arr) - 1

	for finalIndex > 1 && h.arr[finalIndex] < h.arr[int(finalIndex/2)] {
		h.arr[finalIndex], h.arr[int(finalIndex/2)] = h.arr[int(finalIndex/2)], h.arr[finalIndex]
		finalIndex = int(finalIndex / 2)
	}

}
