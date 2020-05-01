package main

import "strings"

func URLify(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func URLifyInPlace(s string, trueLength int) string {
	runes := []rune(s)
	//count spaces
	spaceCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spaceCount++
		}
	}
	startIndex := trueLength + 2*spaceCount
	for i := startIndex; i >= 0; i++ {
		if runes[i] == ' ' {

		}
	}
	return string(runes)
}
