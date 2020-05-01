package main

import (
	"fmt"
)

// Return the longest common subsequence between two strings

/**
input is two strings
"ABAZD" "BACBAD" => "ABAD"
"AAGTAB" "GXTXAYB" => "GTAB"
"aaaa" "aa" => "aa"



**/
func longestCommonSubstring(firstString, secondString string) string {
	//Create a dictionary
	dictionary := make(map[byte]int)
	for i := 0; i < len(firstString); i++ {

		if _, ok := dictionary[firstString[i]]; ok {
			continue
		}

		for j := 0; j < len(secondString); j++ {
			if firstString[i] == secondString[j] {
				dictionary[firstString[i]] = i
			}
		}

	}
	// We have a dictionary with the first entry of the each char of firstString in seondString

	//for each entry in the dictionary try the following fors
	for i := 0; i < len(firstString); i++ {
		// var biggestSubsequence string
		for j := 0; j < len(secondString); j++ {
		}
	}
	// strings.Contains
	return ""
}

func main() {
	firstString := "ABAZD"
	secondString := "BACBAD"
	longest := longestCommonSubstring(firstString, secondString)
	fmt.Println(longest)

	fmt.Println("vim-go")
}
