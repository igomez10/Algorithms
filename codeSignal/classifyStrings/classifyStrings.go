package main

func main() {}

func classifyStrings(s string) string {
	isGood := false
	isBad := false

	for i := 2; i < len(s); i++ {

		vowelCounter := 0     //max3
		consonantCounter := 0 // max5
		for j := 2; j > 0; j-- {
			if s[i-j] == '?' || isVowel(s[i-j]) {
				vowelCounter++
			}
		}

		if vowelCounter == 3 {
			isBad = true
		}

		for j := 4; j > 0; j-- {
			if s[i-j] == '?' || !isVowel(s[i-j]) {
				vowelCounter++
			}
		}

		if consonantCounter == 5 {
			isBad = true
			return "bad"
		}
	}

	if isBad && isGood {
		return "mixed"
	}

	if isBad {
		return "bad"
	}

	if isGood {
		return "good"
	}

	return "none"
}

func isVowel(b byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}

	for v := range vowels {
		if b == vowels[v] {
			return true
		}
	}
	return false
}
