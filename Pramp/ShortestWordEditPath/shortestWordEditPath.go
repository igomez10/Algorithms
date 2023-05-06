package main

import "fmt"

type QueueItem struct {
	Value   string
	Counter int
	Next    *QueueItem
}

func main() {
	source := "bit"
	target := "dog"
	words := []string{"but", "put", "big", "pot", "pog", "dog", "lot"}

	fmt.Println(shortestPath(source, target, words))
}

func shortestPath(source, target string, words []string) int {
	if len(source) != len(target) {
		return -1
	}
	possibleNodes := map[string]bool{}
	for i := range words {
		possibleNodes[words[i]] = true
	}

	queue := &QueueItem{
		Value:   source,
		Counter: 0,
	}

	visited := map[string]bool{
		source: true,
	}

	tail := queue
	for queue != nil {
		currentWord := queue.Value
		currentChanges := queue.Counter
		if currentWord == target {
			return currentChanges
		}

		visited[currentWord] = true
		for nextWord := range possibleNodes {
			if visited[nextWord] {
				continue
			}

			diff := CountDifferences(currentWord, nextWord)
			if diff == 1 {
				tail.Next = &QueueItem{
					Value:   nextWord,
					Counter: currentChanges + 1,
				}
				tail = tail.Next
			}
		}

		queue = queue.Next
	}

	return -1
}

func CountDifferences(a, b string) int {
	counter := 0
	for i := range a {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter
}
