package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(2))
}

func countAndSay(target int) string {
	res := "1"
	for i := 0; i < target-1; i++ {
		res = generateNextSequence(res)
	}
	return res
}

func generateNextSequence(s string) string {
	newString := ""
	counter := 1
	currNumber := s[0]
	for i := 1; i < len(s); i++ {
		if currNumber == s[i] {
			counter++
		} else {
			newString += strconv.Itoa(counter)
			newString += string(currNumber)
			currNumber = s[i]
			counter = 1
		}
	}
	newString += strconv.Itoa(counter)
	newString += string(currNumber)

	return newString
}
