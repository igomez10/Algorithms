package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func reorderLogFiles(logs []string) []string {

	sort.SliceStable(logs, func(i, j int) bool {
		first := strings.Index(logs[i], " ") + 1
		second := strings.Index(logs[j], " ") + 1

		if isDigitLog(logs[i][first]) {
			return false
		} else if !isDigitLog(logs[i][first]) && isDigitLog(logs[j][second]) {
			return true
		} else if !isDigitLog(logs[i][first]) && !isDigitLog(logs[j][second]) {

			comp := strings.Compare(logs[i][first:], logs[j][second:])

			if comp == 0 {
				return strings.Compare(logs[i], logs[j]) == -1
			}

			return comp == -1

		} else {
			panic("unexpected")
		}

	})

	return logs
}

func isDigitLog(b byte) bool {
	s := string(b)
	if _, err := strconv.Atoi(string(s)); err != nil {
		return false
	}

	return true
}

func main() {
	arr := []string{
		"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero",
	}

	fmt.Println(reorderLogFiles(arr))

	// ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]

}
