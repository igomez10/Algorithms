package main

import (
	"fmt"
	"strconv"
	"strings"
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

	newString := strings.Builder{}
	counter := 1
	currNumber := s[0]
	for i := 1; i < len(s); i++ {
		if currNumber == s[i] {
			counter++
		} else {
			newString.WriteString(strconv.Itoa(counter))
			newString.WriteString(string(currNumber))
			currNumber = s[i]
			counter = 1
		}
	}
	newString.WriteString(strconv.Itoa(counter))
	newString.WriteString(string(currNumber))

	return newString.String()
}
