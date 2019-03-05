package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	Name   string
	Expiry int
	Index  int
}

type heapInt []int

func (pq heapInt) Len() int {
	return len(pq)
}

func (pq heapInt) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq *heapInt) Pop() interface{} {
	popped := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return popped
}

func (pq *heapInt) Push(newElement interface{}) {
	newItem := newElement.(int)
	*pq = append(*pq, newItem)
}

func (pq *heapInt) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func main() {
	h := &heapInt{}
	heap.Init(h)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cases := scanner.Text()

	numCases, err := strconv.Atoi(cases)
	if err != nil {
		fmt.Println(cases, err)
	}

	for i := 0; i < numCases; i++ {
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")
		if arr[0] == "1" || arr[0] == "2" {
			first, second := arr[0], arr[1]
			secondNum, err := strconv.Atoi(second)
			if err != nil {
				fmt.Println(secondNum, err)
			}
			if first == "1" {
				heap.Push(h, secondNum)
			} else {
				for i := range *h {
					if (*h)[i] == secondNum {
						heap.Remove(h, i)
						break
					}
				}
			}
		} else {
			fmt.Println((*h)[0])
		}
	}
}
