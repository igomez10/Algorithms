package main

import (
	"fmt"
	"math"
)

func GetShortestUniqueSubstring(arr []string, str string) string {
	dict := map[string]bool{}
	currentCount := map[string]int{}

	for i := range arr {
		dict[string(arr[i])] = true
	}

	lo := 0
	hi := 0
	shortest := math.MaxInt32
	shortestString := ""
	// currentCount[string(str[0])] = 1
	for lo < len(str) {
		if len(currentCount) != len(dict) && hi < len(str) {
			newChar := string(str[hi])
			if dict[newChar] {
				currentCount[newChar]++
			}
			hi++
		} else {
			if len(currentCount) == len(dict) && shortest > hi-lo {
				shortest = hi - lo
				shortestString = str[lo:hi]
			}
			oldChar := string(str[lo])
			if dict[oldChar] {
				currentCount[oldChar]--
			}

			if currentCount[oldChar] == 0 {
				delete(currentCount, oldChar)
			}
			lo++
		}
	}

	return shortestString
}

func main() {
	arr := []string{"x", "y", "z"}
	str := "xyyzyzyx"
	fmt.Println(GetShortestUniqueSubstring(arr, str))
}

/*
['x','y','z']
map {x:0, y: 1, z:1}
counter = len(map)--
//x=0, y=0,
//count--
//


"xyyzyzyx"
map {x: 2, y: 4, z:2}
if s[i] == 'x'{
  map['x']--
  count--
}

x=0
count=0



0 x
map {x:1, y: 0, z:0}

2 y
map {x:1, y: 1, z:0}

*/
