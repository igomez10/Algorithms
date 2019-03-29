package main

import "fmt"

func main() {
	s := "aaaabaaa"
	fmt.Println(findLongestPalindromicSubstring(s))
}

func findLongestPalindromicSubstring(s string) string {
	longest := string(s[0])
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i:j+1]) && j+1-i > len(longest) {
				longest = s[i : j+1]
			}
		}
	}
	return longest
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
