package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	target := 1211
	fmt.Println(countAndSay(target))
}

func countAndSay(soriginal int) string {
	s := strconv.Itoa(soriginal)
	newString := strings.Builder{}
	counter := 1
	currNumber := s[0]
	for i := 1; i < len(s); i++ {
		if currNumber == s[i] {
			counter++
		} else {
			newString.WriteString(strconv.Itoa(counter))
			newString.WriteByte(currNumber)
			currNumber = s[i]
			counter = 1
		}
	}
	newString.WriteString(strconv.Itoa(counter))
	newString.WriteByte(currNumber)

	return newString.String()
}
