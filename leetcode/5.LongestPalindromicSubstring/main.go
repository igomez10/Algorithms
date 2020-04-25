package main

import "fmt"

func main() {
	ss := [][]string{
		[]string{"cbbd", "bb"},
		[]string{"babad", "aba"},
		[]string{"a", "a"},
	}

	for i := range ss {
		if longestPalindrome(ss[i][0]) != ss[i][1] {
			fmt.Printf("FAIL %q %q\n", longestPalindrome(ss[i][0]), ss[i][1])
		} else {
			fmt.Printf("PASS %q %q\n", longestPalindrome(ss[i][0]), ss[i][1])
		}

	}
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	x, y := 0, 0

	for i := 0; i < len(s)-1; i++ {
		if a, b := checkPalindrome(s, i, i); (b - a) >= (y - x) { // odd length palin
			x, y = a, b
		}
		if a, b := checkPalindrome(s, i, i+1); (b - a) >= (y - x) { // even length palin
			x, y = a, b
		}
	}

	return s[x : y+1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkPalindrome(s string, l, r int) (int, int) {
	x, y := 0, 0

	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			break
		}

		x, y = l, r
		l--
		r++
	}

	return x, y
}
