package main

func longestPalindrome(s string) string {
	lo := 0
	hi := 0
	maxPal := string(s[0])

	for lo < len(s) {
		c := expand(s, lo, hi)
		if len(c) > len(maxPal) {
			maxPal = c
		}

		hi++

		c = expand(s, lo, hi)
		if len(c) > len(maxPal) {
			maxPal = c
		}

		lo = hi

	}

	return maxPal
}

func expand(s string, lo, hi int) string {
	if lo < 0 || hi >= len(s) || s[lo] != s[hi] {
		return ""
	}

	for lo > 0 && hi < len(s)-1 && s[lo-1] == s[hi+1] {
		lo--
		hi++
	}

	return string(s[lo : hi+1])
}
