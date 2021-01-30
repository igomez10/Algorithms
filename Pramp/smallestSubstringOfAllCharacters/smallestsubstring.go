package main

import (
	"fmt"
	"math"
)

func GetShortestUniqueSubstring(arr []string, str string) string {
	shortest := math.MaxInt32
	shortestString := ""
	// create a map with the important characters to access them in O(1)
	wanted := map[string]bool{}
	for i := range arr {
		wanted[arr[i]] = true
	}

	// create map of current unique chars in sliding window lo:hi
	uniques := map[string]int{}

	// start a sliding window with lo:hi. Continue until lo == len(str)
	lo := 0
	hi := 0
	for lo < len(str) {

		// if we have missing characters in the current sliding window, extend upper boundary
		if len(uniques) < len(wanted) && hi < len(str) {
			if wanted[string(str[hi])] {
				uniques[string(str[hi])]++ // add new character to unique count
			}
			hi++
		} else {
			// else, we have all the unique characters in the current sliding window, try
			// to increase lo to reduce the length of the sliding window
			if wanted[string(str[lo])] {
				uniques[string(str[lo])]-- // substract old character from unique count
				if uniques[string(str[lo])] == 0 {
					delete(uniques, string(str[lo])) // if reaches 0, remove from unique map
				}
			}
			lo++
		}

		if len(uniques) == len(wanted) {
			if hi-lo < shortest {
				shortest = hi - lo
				shortestString = str[lo:hi]
			}
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
