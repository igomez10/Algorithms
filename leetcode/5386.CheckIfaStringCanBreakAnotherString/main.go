package main

import (
	"fmt"
	"sort"
)

func checkIfCanBreak(string1 string, string2 string) bool {
	s1 := []byte(string1)
	s2 := []byte(string2)

	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })

	fmt.Println(string(s1))
	fmt.Println(string(s2))

	// check if all items in s1 are bigger than s2
	canBreak := true
	for i := range s1 {
		if s1[i] < s2[i] {
			canBreak = false
			break
		}
	}

	if canBreak {
		return true
	} else {
		canBreak = true
		for i := range s2 {
			if s2[i] < s1[i] {
				canBreak = false
				return false
			}
		}
		if canBreak {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("START")
	fmt.Println()

	s1 := "leetcodee"
	s2 := "interview"
	ans := checkIfCanBreak(s1, s2)
	fmt.Println(s1, s2, ans)

	// Input: s1 = "abc", s2 = "xya"
	// Output: true
	s1 = "abc"
	s2 = "xya"
	ans = checkIfCanBreak(s1, s2)
	fmt.Println(s1, s2, ans)

	// Input: s1 = "abe", s2 = "acd"
	// Output: false
	s1 = "abe"
	s2 = "acd"
	ans = checkIfCanBreak(s1, s2)
	fmt.Println(s1, s2, ans)
}
