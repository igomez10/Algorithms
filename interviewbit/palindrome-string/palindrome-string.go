package main

import "fmt"
import "strings"

// Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

func main() {
	val1 := ".,"
	//val1 := "A man, a plan, a canal: Panama"
	//val1 := "1a2"
	fmt.Println(isPalindrome(val1))
}

func isPalindrome(val string) int {

	s := strings.ToLower(val)
	lo := 0
	hi := len(s) - 1
	for lo < hi {
		for lo < len(s) && !isAlphaNumeric(s[lo]) {
			lo++
		}

		for hi >= 0 && !isAlphaNumeric(s[hi]) {
			hi--
		}
		if lo >= hi {
			break
		}
		if s[lo] != s[hi] {
			return 0
		}
		lo++
		hi--
	}
	return 1
}

func isAlphaNumeric(a byte) bool {
	isLetter := a >= 'a' && a <= 'z'
	isNumber := a >= '0' && a <= '9'
	if !isLetter && !isNumber {
		return false
	}
	return true
}
