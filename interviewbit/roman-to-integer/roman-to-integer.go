package main

import "fmt"

/*
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.

Read more details about roman numerals at Roman Numeric System
*/

func romanToInt(romanNum string) int {

	rules := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	counter := 0
	for i := 0; i < len(romanNum)-1; i++ {
		currentNum := rules[rune(romanNum[i])]
		nextNum := rules[rune(romanNum[i+1])]
		if nextNum > currentNum {
			counter -= currentNum
		} else {
			counter += currentNum
		}
	}
	counter += rules[rune(romanNum[len(romanNum)-1])]

	return counter
}

func main() {
	fmt.Println(romanToInt("XIV"))
}
