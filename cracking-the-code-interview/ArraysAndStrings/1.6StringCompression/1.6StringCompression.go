package main

import "strconv"
import "strings"

func stringCompression(s string) string {

	builder := strings.Builder{}

	currCounter := 0
	currChar := s[0]

	for i := range s {
		if s[i] == currChar {
			currCounter++
		} else {
			builder.WriteByte(currChar)
			frequency := strconv.Itoa(currCounter)
			builder.WriteString(frequency)
			currChar = s[i]
			currCounter = 1
		}
	}

	builder.WriteByte(currChar)
	frequency := strconv.Itoa(currCounter)
	builder.WriteString(frequency)
	return builder.String()
}
