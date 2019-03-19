package main

import "fmt"
import "regexp"

func main() {
	value := "abc"
	fmt.Println(validNumber(value))
}

func validNumber(value string) int {
	isMatch, err := regexp.Match(`^[ ]*[-]*[0-9]+([.]?[0-9]+)?[e]?[-]?[0-9]*$`, []byte(value))
	if err != nil {
		fmt.Println(err)
	}
	if isMatch {
		return 1
	}
	return 0
}
