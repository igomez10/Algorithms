package main

import (
	"fmt"
	"math"
)

func main() {
	initialString := "asdcbsdcagfsdbgdfanfghbsfdab"
	answer := alternate(initialString)
	fmt.Println("Answer:", answer)
}
func alternate(initialString string) int32 {
	charsInString := make(map[string]int)

	for _, v := range initialString {
		charsInString[string(v)]++
	}

	possiblePairs := make(map[string][]string)
	for k, v := range charsInString {
		for k2, v2 := range charsInString {
			if k == k2 {
				continue
			} else if math.Abs(float64(v-v2)) <= 1 {
				if _, ok := possiblePairs[k]; ok {
					possiblePairs[k] = append(possiblePairs[k], k2)
				} else {
					possiblePairs[k2] = append(possiblePairs[k2], k)
				}
			}
		}
	}
	var longestString int32
	for k := range possiblePairs {
		for ik2 := range possiblePairs[k] {
			lookingForK := true
			currentStringLength := int32(0)
			for i := range initialString {
				if string(initialString[i]) == k {
					if !lookingForK {
						currentStringLength = 0
						break
					} else {
						currentStringLength++
						lookingForK = !lookingForK
					}
				} else if string(initialString[i]) == possiblePairs[k][ik2] {
					if lookingForK {
						currentStringLength = 0
						break
					} else {
						currentStringLength++
						lookingForK = !lookingForK
					}
				}
			}
			if currentStringLength > longestString {
				longestString = currentStringLength
			}
		}
	}
	return longestString
}
