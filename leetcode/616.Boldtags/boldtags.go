package main

import "sort"

type Node struct {
	Val      string
	Children map[string]*Node
	isLeaf   bool
}

func addBoldTag(sourceString string, dict []string) string {

	root := buildTrie(dict)

	intervals := [][]int{}
	for i := range sourceString {
		found := findWord(root, sourceString, i)
		if len(found) > 1 {
			intervals = append(intervals, found)
		}
	}

	if len(intervals) == 0 {
		return sourceString
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := [][]int{}
	current := intervals[0]
	for i := range intervals {
		if intervals[i][0]-1 <= current[1] {
			if intervals[i][1] > current[1] {
				current[1] = intervals[i][1]
			}
		} else {
			mergedIntervals = append(mergedIntervals, current)
			current = intervals[i]
		}
	}

	mergedIntervals = append(mergedIntervals, current)

	buffer := []byte{}

	for i := range sourceString {
		if len(mergedIntervals) > 0 {
			start := mergedIntervals[0][0]
			end := mergedIntervals[0][1]

			if i == start {
				buffer = append(buffer, []byte("<b>")...)
			}

			buffer = append(buffer, sourceString[i])

			if i == end {
				buffer = append(buffer, []byte("</b>")...)
				mergedIntervals = mergedIntervals[1:]
			}

		} else {
			buffer = append(buffer, sourceString[i])
		}

	}

	return string(buffer)
}

func findWord(root *Node, s string, index int) []int {
	current := root
	interval := []int{index}
	for i := index; i < len(s); i++ {
		currentChar := string(s[i])
		child := current.Children[currentChar]
		if child != nil {
			current = child
		} else {
			break
		}

		if current.isLeaf {
			if len(interval) == 1 {
				interval = append(interval, i)
			} else {
				interval[1] = i
			}
		}
	}

	return interval
}

func buildTrie(dict []string) *Node {
	root := &Node{
		Val:      "",
		Children: map[string]*Node{},
	}

	for i := range dict {
		currentNode := root
		currentWord := dict[i]
		for j := range currentWord {
			currentChar := string(currentWord[j])
			_, exist := currentNode.Children[currentChar]
			if !exist {
				newNode := &Node{
					Val:      currentChar,
					Children: map[string]*Node{},
					isLeaf:   false,
				}
				currentNode.Children[currentChar] = newNode
			}
			currentNode = currentNode.Children[currentChar]
		}
		currentNode.isLeaf = true
	}
	return root
}
