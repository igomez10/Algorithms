package main

import "fmt"

/*
You are given a string S, and you have to find all the amazing substrings of S.

Amazing Substring is one that starts with a vowel (a, e, i, o, u, A, E, I, O, U).


*/

var vowels map[byte]bool

func main() {
	s := "ABEC"
	fmt.Println(numAmazingSubstring(s))
}

func numAmazingSubstring(s string) int {
	vowels = map[byte]bool{
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	counter := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			counter += len(s) - i
		}
	}
	return counter % 10003
}

func isVowel(r byte) bool {
	if _, ok := vowels[r]; ok {
		return true
	}
	return false
}
