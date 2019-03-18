package main

import "fmt"

/*
Write a function to find the longest common prefix
string amongst an array of strings.

Longest common prefix for a pair of strings S1 and S2
is the longest string S which is the prefix of both S1 and S2.

As an example, longest common prefix of "abcdefgh" and "abcefgh" is "abc".

Given the array of strings, you need to find the
longest S which is the prefix of ALL the strings in the array.
*/

func main() {
	arr := []string{
		"abcd",
		"aze",
	}

	fmt.Println(longestCommonPrefix(arr))
}

func longestCommonPrefix(arr []string) string {
	currentLongestPrefix := findSmallest(arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]) && j < len(currentLongestPrefix); j++ {
			if arr[i][j] != currentLongestPrefix[j] {
				currentLongestPrefix = currentLongestPrefix[:j]
				break
			}
		}
	}
	return currentLongestPrefix
}

func findSmallest(arr []string) string {
	var smallest string
	smallest = arr[0]
	for _, v := range arr {
		if len(v) < len(smallest) {
			smallest = v
		}
	}
	return smallest
}
