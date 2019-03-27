package main

func isPermutationOfPalindrome(s string) bool {

	counters := [27]int{}

	for i := range s {
		counters[s[i]-'a']++
	}

	isMiddle := false
	for i := range counters {
		if counters[i]%2 > 0 {
			if !isMiddle {
				isMiddle = true
			} else {
				return false
			}
		}
	}
	return true
}
